package provider

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Repository interface {
	GetAll(context.Context) ([]Provider, error)
	Upsert(context.Context, Provider) (Provider, error)
	GetByID(context.Context, uuid.UUID) (Provider, error)
	GetByName(context.Context, string) (Provider, error)
	DeleteByID(context.Context, uuid.UUID) error
}

var _ Repository = &Service{}

type Service struct {
	providerRepo Repository
}

func NewService(providerRepo Repository) *Service {
	return &Service{
		providerRepo: providerRepo,
	}
}

func (s *Service) DeleteByID(ctx context.Context, providerID uuid.UUID) error {
	return s.providerRepo.DeleteByID(ctx, providerID)
}

func (s *Service) GetAll(ctx context.Context) ([]Provider, error) {
	providers, err := s.providerRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return providers, nil
}

func (s *Service) GetByID(ctx context.Context, providerID uuid.UUID) (Provider, error) {
	provider, err := s.providerRepo.GetByID(ctx, providerID)
	if err != nil {
		return Provider{}, err
	}

	return provider, err
}

func (s *Service) GetByName(ctx context.Context, name string) (Provider, error) {
	provider, err := s.providerRepo.GetByName(ctx, name)
	if err != nil {
		return Provider{}, err
	}

	return provider, err
}

func (s *Service) Upsert(ctx context.Context, c Provider) (Provider, error) {
	id, err := s.providerRepo.Upsert(ctx, c)
	if err != nil {
		return Provider{}, fmt.Errorf("failed to save provider: %w", err)
	}

	return id, err
}
