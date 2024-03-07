package postgres

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/missingstudio/ai/gateway/core/prompt"
)

type PromptDB struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Template    string    `db:"template"`
	Metadata    []byte    `db:"metadata"`
	UpdatedAt   time.Time `db:"updated_at"`
	CreatedAt   time.Time `db:"created_at"`
}

func (c PromptDB) ToPrompt() (prompt.Prompt, error) {
	var unmarshalledMetadata map[string]any
	if len(c.Metadata) > 0 {
		if err := json.Unmarshal(c.Metadata, &unmarshalledMetadata); err != nil {
			return prompt.Prompt{}, fmt.Errorf("failed to unmarshal connection metadata(%s): %w", c.ID.String(), err)
		}
	}

	return prompt.Prompt{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		Template:    c.Template,
		Metadata:    unmarshalledMetadata,
	}, nil
}
