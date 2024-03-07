package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/missingstudio/ai/gateway/core/connection"
	"github.com/missingstudio/ai/gateway/pkg/database"
)

var _ connection.Repository = &ConnectionRepository{}

var (
	ErrConflict = errors.New("connection already exist")
	ErrNotExist = errors.New("apikey or its relations doesn't exist")
)

type ConnectionRepository struct {
	dbc *database.Client
}

func NewConnectionRepository(dbc *database.Client) *ConnectionRepository {
	return &ConnectionRepository{
		dbc: dbc,
	}
}

func (*ConnectionRepository) DeleteByID(ctx context.Context, connID uuid.UUID) error {
	panic("unimplemented")
}

func (c *ConnectionRepository) GetAll(ctx context.Context) ([]connection.Connection, error) {
	query, params, err := dialect.From(TABLE_CONNECTIONS).ToSQL()
	if err != nil {
		return []connection.Connection{}, fmt.Errorf("%w: %s", queryErr, err)
	}

	var conns []ConnectionDB
	if err = c.dbc.WithTimeout(ctx, TABLE_CONNECTIONS, "List", func(ctx context.Context) error {
		return c.dbc.SelectContext(ctx, &conns, query, params...)
	}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []connection.Connection{}, fmt.Errorf("%s", err)
		}
		return []connection.Connection{}, fmt.Errorf("%w: %s", dbErr, err)
	}

	var connections []connection.Connection
	for _, c := range conns {
		conn, err := c.ToConnection()
		if err != nil {
			return []connection.Connection{}, fmt.Errorf("%w: %s", parseErr, err)
		}

		connections = append(connections, conn)
	}

	return connections, nil
}

func (*ConnectionRepository) GetByID(ctx context.Context, connID uuid.UUID) (connection.Connection, error) {
	panic("unimplemented")
}

func (c *ConnectionRepository) GetByName(ctx context.Context, name string) (connection.Connection, error) {
	query, params, err := dialect.From(TABLE_CONNECTIONS).Where(goqu.Ex{"name": name}).ToSQL()
	if err != nil {
		return connection.Connection{}, err
	}

	// Execute the SQL query and retrieve the result
	var connDb ConnectionDB
	if err = c.dbc.WithTimeout(ctx, TABLE_CONNECTIONS, "Get", func(ctx context.Context) error {
		return c.dbc.QueryRowxContext(ctx, query, params...).StructScan(&connDb)
	}); err != nil {
		err = checkPostgresError(err)
		switch {
		case errors.Is(err, ErrDuplicateKey):
			return connection.Connection{}, ErrConflict
		default:
			return connection.Connection{}, err
		}
	}

	return connDb.ToConnection()
}

func (c *ConnectionRepository) Upsert(ctx context.Context, conn connection.Connection) (connection.Connection, error) {
	marshaledConfig, err := json.Marshal(conn.Config)
	if err != nil {
		return connection.Connection{}, fmt.Errorf("namespace metadata: %w: %s", parseErr, err)
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
		return connection.Connection{}, fmt.Errorf("%w: %s", queryErr, err)
	}

	var connDb ConnectionDB
	if err = c.dbc.WithTimeout(ctx, TABLE_CONNECTIONS, "Upsert", func(ctx context.Context) error {
		return c.dbc.QueryRowxContext(ctx, query, params...).StructScan(&connDb)
	}); err != nil {
		err = checkPostgresError(err)
		switch {
		case errors.Is(err, ErrDuplicateKey):
			return connection.Connection{}, ErrConflict
		default:
			return connection.Connection{}, err
		}
	}

	return connDb.ToConnection()
}
