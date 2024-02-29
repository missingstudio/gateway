package apikey

import (
	"time"

	"github.com/google/uuid"
)

type APIKey struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Value       string    `json:"value"`
	MaskedValue string    `json:"masked_value"`
	CreatedAt   time.Time `json:"created_at"`
	LastUsedAt  time.Time `json:"last_used_at"`
}
