package postgres

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/missingstudio/ai/gateway/core/provider"
)

type ProviderDB struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Config    []byte    `db:"config"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

func (c ProviderDB) ToProvider() (provider.Provider, error) {
	var unmarshalledConfig map[string]any
	if len(c.Config) > 0 {
		if err := json.Unmarshal(c.Config, &unmarshalledConfig); err != nil {
			return provider.Provider{}, fmt.Errorf("failed to unmarshal connection config(%s): %w", c.ID.String(), err)
		}
	}

	return provider.Provider{
		ID:     c.ID,
		Name:   c.Name,
		Config: unmarshalledConfig,
	}, nil
}
