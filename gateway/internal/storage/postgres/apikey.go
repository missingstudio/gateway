package postgres

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/missingstudio/ai/gateway/core/apikey"
)

type APIKey struct {
	ID         uuid.UUID    `db:"id"`
	Name       string       `db:"name"`
	Value      []byte       `db:"value"`
	LastUsedAt sql.NullTime `db:"last_used_at"`
	UpdatedAt  time.Time    `db:"updated_at"`
	CreatedAt  time.Time    `db:"created_at"`
}

func (c APIKey) ToAPIKey() (apikey.APIKey, error) {
	return apikey.APIKey{
		Id:         c.ID,
		Name:       c.Name,
		Value:      string(c.Value),
		LastUsedAt: c.LastUsedAt.Time,
		CreatedAt:  c.CreatedAt,
	}, nil
}
