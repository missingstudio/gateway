package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/missingstudio/studio/backend/core/apikey"
	"github.com/missingstudio/studio/backend/pkg/database"
)

var _ apikey.Repository = &APIKeyRepository{}

type APIKeyRepository struct {
	dbc *database.Client
}

func NewAPIKeyRepository(dbc *database.Client) *APIKeyRepository {
	return &APIKeyRepository{
		dbc: dbc,
	}
}

func (c *APIKeyRepository) GetAll(ctx context.Context) ([]apikey.APIKey, error) {
	query, params, err := dialect.From(TABLE_APIKEYS).ToSQL()
	if err != nil {
		return []apikey.APIKey{}, fmt.Errorf("%w: %s", queryErr, err)
	}

	var apis []APIKey
	if err = c.dbc.WithTimeout(ctx, TABLE_APIKEYS, "List", func(ctx context.Context) error {
		return c.dbc.SelectContext(ctx, &apis, query, params...)
	}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []apikey.APIKey{}, fmt.Errorf("%s", err)
		}
		return []apikey.APIKey{}, fmt.Errorf("%w: %s", dbErr, err)
	}

	var mkeys []apikey.APIKey
	for _, c := range apis {
		prompt, err := c.ToAPIKey()
		if err != nil {
			return []apikey.APIKey{}, fmt.Errorf("%w: %s", parseErr, err)
		}

		mkeys = append(mkeys, prompt)
	}

	return mkeys, nil
}

func (ak *APIKeyRepository) Create(ctx context.Context, key apikey.APIKey) (apikey.APIKey, error) {
	query, params, err := dialect.Insert(TABLE_APIKEYS).Rows(
		goqu.Record{
			"name":  key.Name,
			"value": key.Value,
		}).Returning(&APIKey{}).ToSQL()
	if err != nil {
		return apikey.APIKey{}, fmt.Errorf("%w: %s", queryErr, err)
	}

	var akey APIKey
	if err = ak.dbc.WithTimeout(ctx, TABLE_APIKEYS, "Upsert", func(ctx context.Context) error {
		return ak.dbc.QueryRowxContext(ctx, query, params...).StructScan(&akey)
	}); err != nil {
		err = checkPostgresError(err)
		switch {
		case errors.Is(err, ErrDuplicateKey):
			return apikey.APIKey{}, ErrConflict
		default:
			return apikey.APIKey{}, err
		}
	}

	return akey.ToAPIKey()
}

func (c *APIKeyRepository) Get(ctx context.Context, id string) (apikey.APIKey, error) {
	query, params, err := dialect.From(TABLE_APIKEYS).Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		return apikey.APIKey{}, err
	}

	var akey APIKey
	if err = c.dbc.WithTimeout(ctx, TABLE_APIKEYS, "Get", func(ctx context.Context) error {
		return c.dbc.QueryRowxContext(ctx, query, params...).StructScan(&akey)
	}); err != nil {
		err = checkPostgresError(err)
		switch {
		case errors.Is(err, ErrDuplicateKey):
			return apikey.APIKey{}, ErrConflict
		default:
			return apikey.APIKey{}, err
		}
	}

	return akey.ToAPIKey()
}

func (c *APIKeyRepository) Update(ctx context.Context, key apikey.APIKey) (apikey.APIKey, error) {
	query, params, err := dialect.Insert(TABLE_APIKEYS).Rows(
		goqu.Record{
			"id":    key.Id,
			"name":  key.Name,
			"value": key.Value,
		}).OnConflict(
		goqu.DoUpdate("id", goqu.Record{
			"name": key.Name,
		})).Returning(&APIKey{}).ToSQL()
	if err != nil {
		return apikey.APIKey{}, fmt.Errorf("%w: %s", queryErr, err)
	}

	var akey APIKey
	if err = c.dbc.WithTimeout(ctx, TABLE_APIKEYS, "Upsert", func(ctx context.Context) error {
		return c.dbc.QueryRowxContext(ctx, query, params...).StructScan(&akey)
	}); err != nil {
		err = checkPostgresError(err)
		switch {
		case errors.Is(err, ErrDuplicateKey):
			return apikey.APIKey{}, ErrConflict
		default:
			return apikey.APIKey{}, err
		}
	}

	return akey.ToAPIKey()
}

func (c *APIKeyRepository) DeleteByID(ctx context.Context, id string) error {
	query, params, err := dialect.Delete(TABLE_APIKEYS).Where(
		goqu.Ex{
			"id": id,
		},
	).ToSQL()
	if err != nil {
		return fmt.Errorf("%w: %s", queryErr, err)
	}

	if err = c.dbc.WithTimeout(ctx, TABLE_APIKEYS, "Delete", func(ctx context.Context) error {
		if _, err = c.dbc.DB.ExecContext(ctx, query, params...); err != nil {
			return err
		}
		return nil
	}); err != nil {
		err = checkPostgresError(err)
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrNotExist
		default:
			return err
		}
	}
	return nil
}

func (c *APIKeyRepository) GetByToken(ctx context.Context, token string) (apikey.APIKey, error) {
	query, params, err := dialect.From(TABLE_APIKEYS).Where(goqu.Ex{"value": token}).ToSQL()
	if err != nil {
		return apikey.APIKey{}, err
	}

	var akey APIKey
	if err = c.dbc.WithTimeout(ctx, TABLE_APIKEYS, "Get", func(ctx context.Context) error {
		return c.dbc.QueryRowxContext(ctx, query, params...).StructScan(&akey)
	}); err != nil {
		err = checkPostgresError(err)
		switch {
		case errors.Is(err, ErrDuplicateKey):
			return apikey.APIKey{}, ErrConflict
		default:
			return apikey.APIKey{}, err
		}
	}

	return akey.ToAPIKey()
}
