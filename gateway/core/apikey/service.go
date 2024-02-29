package apikey

import (
	"context"
)

type Repository interface {
	GetAll(context.Context) ([]APIKey, error)
	Create(context.Context, APIKey) (APIKey, error)
	Get(context.Context, string) (APIKey, error)
	GetByToken(context.Context, string) (APIKey, error)
	Update(context.Context, APIKey) (APIKey, error)
	DeleteByID(context.Context, string) error
}

var _ Repository = &Service{}

type Service struct {
	apikeyRepo Repository
}

func NewService(apikeyRepo Repository) *Service {
	return &Service{
		apikeyRepo: apikeyRepo,
	}
}

func (s *Service) GetAll(ctx context.Context) ([]APIKey, error) {
	return s.apikeyRepo.GetAll(ctx)
}

func (s *Service) Create(ctx context.Context, api APIKey) (APIKey, error) {
	return s.apikeyRepo.Create(ctx, api)
}

func (s *Service) Get(ctx context.Context, id string) (APIKey, error) {
	return s.apikeyRepo.Get(ctx, id)
}

func (s *Service) GetByToken(ctx context.Context, id string) (APIKey, error) {
	return s.apikeyRepo.GetByToken(ctx, id)
}

func (s *Service) Update(ctx context.Context, api APIKey) (APIKey, error) {
	return s.apikeyRepo.Update(ctx, api)
}

func (s *Service) DeleteByID(ctx context.Context, id string) error {
	return s.apikeyRepo.DeleteByID(ctx, id)
}
