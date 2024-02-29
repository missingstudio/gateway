package connection

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Repository interface {
	GetAll(ctx context.Context) ([]Connection, error)
	Upsert(ctx context.Context, c Connection) (Connection, error)
	GetByID(ctx context.Context, connID uuid.UUID) (Connection, error)
	GetByName(ctx context.Context, name string) (Connection, error)
	DeleteByID(ctx context.Context, connID uuid.UUID) error
}

var _ Repository = &Service{}

type Service struct {
	connectionRepo Repository
}

func NewService(connectionRepo Repository) *Service {
	return &Service{
		connectionRepo: connectionRepo,
	}
}

func (s *Service) DeleteByID(ctx context.Context, connID uuid.UUID) error {
	return s.connectionRepo.DeleteByID(ctx, connID)
}

func (s *Service) GetAll(ctx context.Context) ([]Connection, error) {
	conns, err := s.connectionRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return conns, nil
}

func (s *Service) GetByID(ctx context.Context, connID uuid.UUID) (Connection, error) {
	conn, err := s.connectionRepo.GetByID(ctx, connID)
	if err != nil {
		return Connection{}, err
	}

	return conn, err
}

func (s *Service) GetByName(ctx context.Context, name string) (Connection, error) {
	conn, err := s.connectionRepo.GetByName(ctx, name)
	if err != nil {
		return Connection{}, err
	}

	return conn, err
}

func (s *Service) Upsert(ctx context.Context, c Connection) (Connection, error) {
	id, err := s.connectionRepo.Upsert(ctx, c)
	if err != nil {
		return Connection{}, fmt.Errorf("failed to save connection: %w", err)
	}

	return id, err
}
