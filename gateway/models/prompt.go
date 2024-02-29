package models

import "github.com/google/uuid"

type PromptState string

const (
	PromptStateActive   PromptState = "active"
	PromptStateInactive PromptState = "inactive"
)

type Prompt struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Template    string         `json:"template"`
	State       PromptState    `json:"state"`
	Metadata    map[string]any `json:"metadata"`
}
