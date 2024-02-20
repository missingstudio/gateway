package connections

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/missingstudio/studio/backend/models"
)

type Repository interface {
	GetAll(ctx context.Context) ([]models.Connection, error)
	Upsert(ctx context.Context, c models.Connection) (*models.Connection, error)
	GetByID(ctx context.Context, connID uuid.UUID) (*models.Connection, error)
	GetByName(ctx context.Context, name string) (*models.Connection, error)
	DeleteByID(ctx context.Context, connID uuid.UUID) error
}
