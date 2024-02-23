package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/gofrs/uuid"
	"github.com/missingstudio/studio/backend/internal/connections"
	"github.com/missingstudio/studio/backend/models"
	"github.com/missingstudio/studio/backend/pkg/database"
)

var _ connections.Repository = &ConnectionRepository{}

var ErrConflict = errors.New("connection already exist")

type ConnectionRepository struct {
	dbc *database.Client
}

func NewConnectionRepository(dbc *database.Client) *ConnectionRepository {
	return &ConnectionRepository{
		dbc: dbc,
	}
}

// DeleteByID implements connection.Repository.
func (*ConnectionRepository) DeleteByID(ctx context.Context, connID uuid.UUID) error {
	panic("unimplemented")
}

// GetAll implements connection.Repository.
func (c *ConnectionRepository) GetAll(ctx context.Context) ([]models.Connection, error) {
	query, params, err := dialect.From(TABLE_CONNECTIONS).ToSQL()
	if err != nil {
		return []models.Connection{}, fmt.Errorf("%w: %s", queryErr, err)
	}

	var conns []ConnectionDB
	if err = c.dbc.WithTimeout(ctx, TABLE_CONNECTIONS, "List", func(ctx context.Context) error {
		return c.dbc.SelectContext(ctx, &conns, query, params...)
	}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []models.Connection{}, fmt.Errorf("%s", err)
		}
		return []models.Connection{}, fmt.Errorf("%w: %s", dbErr, err)
	}

	var connections []models.Connection
	for _, c := range conns {
		connection, err := c.ToConnection()
		if err != nil {
			return []models.Connection{}, fmt.Errorf("%w: %s", parseErr, err)
		}

		connections = append(connections, connection)
	}

	return connections, nil
}

// GetByID implements connection.Repository.
func (*ConnectionRepository) GetByID(ctx context.Context, connID uuid.UUID) (models.Connection, error) {
	panic("unimplemented")
}

// GetByID implements connection.Repository.
func (c *ConnectionRepository) GetByName(ctx context.Context, name string) (models.Connection, error) {
	query, params, err := dialect.From(TABLE_CONNECTIONS).Where(goqu.Ex{"name": name}).ToSQL()
	if err != nil {
		return models.Connection{}, err
	}

	// Execute the SQL query and retrieve the result
	var connection ConnectionDB
	if err = c.dbc.WithTimeout(ctx, TABLE_CONNECTIONS, "Get", func(ctx context.Context) error {
		return c.dbc.QueryRowxContext(ctx, query, params...).StructScan(&connection)
	}); err != nil {
		err = checkPostgresError(err)
		switch {
		case errors.Is(err, ErrDuplicateKey):
			return models.Connection{}, ErrConflict
		default:
			return models.Connection{}, err
		}
	}

	return connection.ToConnection()
}

// Upsert implements connection.Repository.
func (c *ConnectionRepository) Upsert(ctx context.Context, conn models.Connection) (models.Connection, error) {
	marshaledConfig, err := json.Marshal(conn.Config)
	if err != nil {
		return models.Connection{}, fmt.Errorf("namespace metadata: %w: %s", parseErr, err)
	}

	query, params, err := dialect.Insert(TABLE_CONNECTIONS).Rows(
		goqu.Record{
			"name":   conn.Name,
			"config": marshaledConfig,
		}).OnConflict(
		goqu.DoUpdate("name", goqu.Record{
			"config":     marshaledConfig,
			"updated_at": goqu.L("now()"),
		})).Returning(&ConnectionDB{}).ToSQL()
	if err != nil {
		return models.Connection{}, fmt.Errorf("%w: %s", queryErr, err)
	}

	var connection ConnectionDB
	if err = c.dbc.WithTimeout(ctx, TABLE_CONNECTIONS, "Upsert", func(ctx context.Context) error {
		return c.dbc.QueryRowxContext(ctx, query, params...).StructScan(&connection)
	}); err != nil {
		err = checkPostgresError(err)
		switch {
		case errors.Is(err, ErrDuplicateKey):
			return models.Connection{}, ErrConflict
		default:
			return models.Connection{}, err
		}
	}

	return connection.ToConnection()
}
