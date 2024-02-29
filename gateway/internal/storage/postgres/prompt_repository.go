package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/missingstudio/studio/backend/internal/prompt"
	"github.com/missingstudio/studio/backend/models"
	"github.com/missingstudio/studio/backend/pkg/database"
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

func (c *PromptRepository) GetAll(ctx context.Context) ([]models.Prompt, error) {
	query, params, err := dialect.From(TABLE_PROMPTS).ToSQL()
	if err != nil {
		return []models.Prompt{}, fmt.Errorf("%w: %s", queryErr, err)
	}

	var pms []PromptDB
	if err = c.dbc.WithTimeout(ctx, TABLE_PROMPTS, "List", func(ctx context.Context) error {
		return c.dbc.SelectContext(ctx, &pms, query, params...)
	}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []models.Prompt{}, fmt.Errorf("%s", err)
		}
		return []models.Prompt{}, fmt.Errorf("%w: %s", dbErr, err)
	}

	var prompts []models.Prompt
	for _, c := range pms {
		prompt, err := c.ToPrompt()
		if err != nil {
			return []models.Prompt{}, fmt.Errorf("%w: %s", parseErr, err)
		}

		prompts = append(prompts, prompt)
	}

	return prompts, nil
}

func (*PromptRepository) GetByID(ctx context.Context, connID uuid.UUID) (models.Prompt, error) {
	panic("unimplemented")
}

func (c *PromptRepository) GetByName(ctx context.Context, name string) (models.Prompt, error) {
	query, params, err := dialect.From(TABLE_PROMPTS).Where(goqu.Ex{"name": name}).ToSQL()
	if err != nil {
		return models.Prompt{}, err
	}

	var prompt PromptDB
	if err = c.dbc.WithTimeout(ctx, TABLE_PROMPTS, "Get", func(ctx context.Context) error {
		return c.dbc.QueryRowxContext(ctx, query, params...).StructScan(&prompt)
	}); err != nil {
		err = checkPostgresError(err)
		switch {
		case errors.Is(err, ErrDuplicateKey):
			return models.Prompt{}, ErrConflict
		default:
			return models.Prompt{}, err
		}
	}

	return prompt.ToPrompt()
}

// Upsert implements prompt.Repository.
func (c *PromptRepository) Upsert(ctx context.Context, conn models.Prompt) (models.Prompt, error) {
	marshaledMetadata, err := json.Marshal(conn.Metadata)
	if err != nil {
		return models.Prompt{}, fmt.Errorf("namespace metadata: %w: %s", parseErr, err)
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
		return models.Prompt{}, fmt.Errorf("%w: %s", queryErr, err)
	}

	var prompt PromptDB
	if err = c.dbc.WithTimeout(ctx, TABLE_PROMPTS, "Upsert", func(ctx context.Context) error {
		return c.dbc.QueryRowxContext(ctx, query, params...).StructScan(&prompt)
	}); err != nil {
		err = checkPostgresError(err)
		switch {
		case errors.Is(err, ErrDuplicateKey):
			return models.Prompt{}, ErrConflict
		default:
			return models.Prompt{}, err
		}
	}

	return prompt.ToPrompt()
}

func (*PromptRepository) DeleteByID(ctx context.Context, connID uuid.UUID) error {
	panic("unimplemented")
}
