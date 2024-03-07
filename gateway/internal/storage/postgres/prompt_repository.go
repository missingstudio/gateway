package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/missingstudio/ai/gateway/core/prompt"
	"github.com/missingstudio/ai/gateway/pkg/database"
)

var _ prompt.Repository = &PromptRepository{}

type PromptRepository struct {
	dbc *database.Client
}

func NewPromptRepository(dbc *database.Client) *PromptRepository {
	return &PromptRepository{
		dbc: dbc,
	}
}

func (c *PromptRepository) GetAll(ctx context.Context) ([]prompt.Prompt, error) {
	query, params, err := dialect.From(TABLE_PROMPTS).ToSQL()
	if err != nil {
		return []prompt.Prompt{}, fmt.Errorf("%w: %s", queryErr, err)
	}

	var pms []PromptDB
	if err = c.dbc.WithTimeout(ctx, TABLE_PROMPTS, "List", func(ctx context.Context) error {
		return c.dbc.SelectContext(ctx, &pms, query, params...)
	}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []prompt.Prompt{}, fmt.Errorf("%s", err)
		}
		return []prompt.Prompt{}, fmt.Errorf("%w: %s", dbErr, err)
	}

	var prompts []prompt.Prompt
	for _, c := range pms {
		p, err := c.ToPrompt()
		if err != nil {
			return []prompt.Prompt{}, fmt.Errorf("%w: %s", parseErr, err)
		}

		prompts = append(prompts, p)
	}

	return prompts, nil
}

func (*PromptRepository) GetByID(ctx context.Context, connID uuid.UUID) (prompt.Prompt, error) {
	panic("unimplemented")
}

func (c *PromptRepository) GetByName(ctx context.Context, name string) (prompt.Prompt, error) {
	query, params, err := dialect.From(TABLE_PROMPTS).Where(goqu.Ex{"name": name}).ToSQL()
	if err != nil {
		return prompt.Prompt{}, err
	}

	var pdb PromptDB
	if err = c.dbc.WithTimeout(ctx, TABLE_PROMPTS, "Get", func(ctx context.Context) error {
		return c.dbc.QueryRowxContext(ctx, query, params...).StructScan(&pdb)
	}); err != nil {
		err = checkPostgresError(err)
		switch {
		case errors.Is(err, ErrDuplicateKey):
			return prompt.Prompt{}, ErrConflict
		default:
			return prompt.Prompt{}, err
		}
	}

	return pdb.ToPrompt()
}

func (c *PromptRepository) Upsert(ctx context.Context, conn prompt.Prompt) (prompt.Prompt, error) {
	marshaledMetadata, err := json.Marshal(conn.Metadata)
	if err != nil {
		return prompt.Prompt{}, fmt.Errorf("namespace metadata: %w: %s", parseErr, err)
	}

	query, params, err := dialect.Insert(TABLE_PROMPTS).Rows(
		goqu.Record{
			"name":     conn.Name,
			"metadata": marshaledMetadata,
		}).OnConflict(
		goqu.DoUpdate("name", goqu.Record{
			"description": conn.Description,
			"template":    conn.Template,
			"metadata":    marshaledMetadata,
			"updated_at":  goqu.L("now()"),
		})).Returning(&PromptDB{}).ToSQL()
	if err != nil {
		return prompt.Prompt{}, fmt.Errorf("%w: %s", queryErr, err)
	}

	var pdb PromptDB
	if err = c.dbc.WithTimeout(ctx, TABLE_PROMPTS, "Upsert", func(ctx context.Context) error {
		return c.dbc.QueryRowxContext(ctx, query, params...).StructScan(&pdb)
	}); err != nil {
		err = checkPostgresError(err)
		switch {
		case errors.Is(err, ErrDuplicateKey):
			return prompt.Prompt{}, ErrConflict
		default:
			return prompt.Prompt{}, err
		}
	}

	return pdb.ToPrompt()
}

func (*PromptRepository) DeleteByID(ctx context.Context, connID uuid.UUID) error {
	panic("unimplemented")
}
