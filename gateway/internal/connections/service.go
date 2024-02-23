package connections

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/missingstudio/studio/backend/models"
)

var _ Repository = &Service{}

type Service struct {
	connectionRepo Repository
}

func NewService(connectionRepo Repository) *Service {
	return &Service{
		connectionRepo: connectionRepo,
	}
}

// DeleteByID implements connection.Repository.
func (s *Service) DeleteByID(ctx context.Context, connID uuid.UUID) error {
	return s.connectionRepo.DeleteByID(ctx, connID)
}

// GetAll implements connection.Repository.
func (s *Service) GetAll(ctx context.Context) ([]models.Connection, error) {
	conns, err := s.connectionRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return conns, nil
}

// GetByID implements connection.Repository.
func (s *Service) GetByID(ctx context.Context, connID uuid.UUID) (models.Connection, error) {
	conn, err := s.connectionRepo.GetByID(ctx, connID)
	if err != nil {
		return models.Connection{}, err
	}

	return conn, err
}

// GetByName implements connection.Repository.
func (s *Service) GetByName(ctx context.Context, name string) (models.Connection, error) {
	conn, err := s.connectionRepo.GetByName(ctx, name)
	if err != nil {
		return models.Connection{}, err
	}

	return conn, err
}

// Upsert implements connection.Repository.
func (s *Service) Upsert(ctx context.Context, c models.Connection) (models.Connection, error) {
	id, err := s.connectionRepo.Upsert(ctx, c)
	if err != nil {
		return models.Connection{}, fmt.Errorf("failed to save connection: %w", err)
	}

	return id, err
}
