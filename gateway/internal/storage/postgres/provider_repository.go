package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/missingstudio/ai/gateway/core/provider"
	"github.com/missingstudio/ai/gateway/pkg/database"
)

var _ provider.Repository = &ProviderRepository{}

var (
	ErrConflict = errors.New("provider already exist")
	ErrNotExist = errors.New("apikey or its relations doesn't exist")
)

type ProviderRepository struct {
	dbc *database.Client
}

func NewProviderRepository(dbc *database.Client) *ProviderRepository {
	return &ProviderRepository{
		dbc: dbc,
	}
}

func (c *ProviderRepository) GetAll(ctx context.Context) ([]provider.Provider, error) {
	query, params, err := dialect.From(TABLE_PROVIDERS).ToSQL()
	if err != nil {
		return []provider.Provider{}, fmt.Errorf("%w: %s", queryErr, err)
	}

	var prDB []ProviderDB
	if err = c.dbc.WithTimeout(ctx, TABLE_PROVIDERS, "List", func(ctx context.Context) error {
		return c.dbc.SelectContext(ctx, &prDB, query, params...)
	}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []provider.Provider{}, fmt.Errorf("%s", err)
		}
		return []provider.Provider{}, fmt.Errorf("%w: %s", dbErr, err)
	}

	var connections []provider.Provider
	for _, c := range prDB {
		pr, err := c.ToProvider()
		if err != nil {
			return []provider.Provider{}, fmt.Errorf("%w: %s", parseErr, err)
		}

		connections = append(connections, pr)
	}

	return connections, nil
}

func (*ProviderRepository) GetByID(ctx context.Context, connID uuid.UUID) (provider.Provider, error) {
	panic("unimplemented")
}

func (c *ProviderRepository) GetByName(ctx context.Context, name string) (provider.Provider, error) {
	query, params, err := dialect.From(TABLE_PROVIDERS).Where(goqu.Ex{"name": name}).ToSQL()
	if err != nil {
		return provider.Provider{}, err
	}

	// Execute the SQL query and retrieve the result
	var prDB ProviderDB
	if err = c.dbc.WithTimeout(ctx, TABLE_PROVIDERS, "Get", func(ctx context.Context) error {
		return c.dbc.QueryRowxContext(ctx, query, params...).StructScan(&prDB)
	}); err != nil {
		err = checkPostgresError(err)
		switch {
		case errors.Is(err, ErrDuplicateKey):
			return provider.Provider{}, ErrConflict
		default:
			return provider.Provider{}, err
		}
	}

	return prDB.ToProvider()
}

func (c *ProviderRepository) Upsert(ctx context.Context, pr provider.Provider) (provider.Provider, error) {
	marshaledConfig, err := json.Marshal(pr.Config)
	if err != nil {
		return provider.Provider{}, fmt.Errorf("namespace metadata: %w: %s", parseErr, err)
	}

	query, params, err := dialect.Insert(TABLE_PROVIDERS).Rows(
		goqu.Record{
			"name":   pr.Name,
			"config": marshaledConfig,
		}).OnConflict(
		goqu.DoUpdate("name", goqu.Record{
			"config":     marshaledConfig,
			"updated_at": goqu.L("now()"),
		})).Returning(&ProviderDB{}).ToSQL()
	if err != nil {
		return provider.Provider{}, fmt.Errorf("%w: %s", queryErr, err)
	}

	var prDB ProviderDB
	if err = c.dbc.WithTimeout(ctx, TABLE_PROVIDERS, "Upsert", func(ctx context.Context) error {
		return c.dbc.QueryRowxContext(ctx, query, params...).StructScan(&prDB)
	}); err != nil {
		err = checkPostgresError(err)
		switch {
		case errors.Is(err, ErrDuplicateKey):
			return provider.Provider{}, ErrConflict
		default:
			return provider.Provider{}, err
		}
	}

	return prDB.ToProvider()
}

func (*ProviderRepository) DeleteByID(ctx context.Context, connID uuid.UUID) error {
	panic("unimplemented")
}
