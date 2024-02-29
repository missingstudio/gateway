package prompt

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Repository interface {
	GetAll(context.Context) ([]Prompt, error)
	Upsert(context.Context, Prompt) (Prompt, error)
	GetByID(context.Context, uuid.UUID) (Prompt, error)
	GetByName(context.Context, string) (Prompt, error)
	DeleteByID(context.Context, uuid.UUID) error
}

var _ Repository = &Service{}

type Service struct {
	promptRepo Repository
}

func NewService(promptRepo Repository) *Service {
	return &Service{
		promptRepo: promptRepo,
	}
}

func (s *Service) DeleteByID(ctx context.Context, promptID uuid.UUID) error {
	return s.promptRepo.DeleteByID(ctx, promptID)
}

func (s *Service) GetAll(ctx context.Context) ([]Prompt, error) {
	prompts, err := s.promptRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return prompts, nil
}

func (s *Service) GetByID(ctx context.Context, promptID uuid.UUID) (Prompt, error) {
	prompt, err := s.promptRepo.GetByID(ctx, promptID)
	if err != nil {
		return Prompt{}, err
	}

	return prompt, err
}

func (s *Service) GetByName(ctx context.Context, name string) (Prompt, error) {
	prompt, err := s.promptRepo.GetByName(ctx, name)
	if err != nil {
		return Prompt{}, err
	}

	return prompt, err
}

func (s *Service) Upsert(ctx context.Context, c Prompt) (Prompt, error) {
	id, err := s.promptRepo.Upsert(ctx, c)
	if err != nil {
		return Prompt{}, fmt.Errorf("failed to save prompt: %w", err)
	}

	return id, err
}
