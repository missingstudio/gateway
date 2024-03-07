package postgres

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/missingstudio/ai/gateway/core/connection"
)

type ConnectionDB struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Config    []byte    `db:"config"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

func (c ConnectionDB) ToConnection() (connection.Connection, error) {
	var unmarshalledConfig map[string]any
	if len(c.Config) > 0 {
		if err := json.Unmarshal(c.Config, &unmarshalledConfig); err != nil {
			return connection.Connection{}, fmt.Errorf("failed to unmarshal connection config(%s): %w", c.ID.String(), err)
		}
	}

	return connection.Connection{
		ID:     c.ID,
		Name:   c.Name,
		Config: unmarshalledConfig,
	}, nil
}
