package tryout

import (
	"context"

	"github.com/doodpanda/tryout-backend/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service interface {
	GetTryoutList(ctx context.Context, param pgtype.UUID) ([]*repository.Tryout, error)
	GetTryoutListFiltered(ctx context.Context, param repository.GetTryoutListFilteredParams) ([]*repository.Tryout, error)
	GetTryoutById(ctx context.Context, tryoutID pgtype.UUID) (*repository.Tryout, error)
}

type service struct {
	repo repository.Querier
}

func NewService(repo repository.Querier) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetTryoutList(ctx context.Context, param pgtype.UUID) ([]*repository.Tryout, error) {
	return s.repo.GetTryoutList(ctx, param)
}

func (s *service) GetTryoutListFiltered(ctx context.Context, param repository.GetTryoutListFilteredParams) ([]*repository.Tryout, error) {
	return s.repo.GetTryoutListFiltered(ctx, &param)
}

func (s *service) GetTryoutById(ctx context.Context, tryoutID pgtype.UUID) (*repository.Tryout, error) {
	return s.repo.GetTryoutById(ctx, tryoutID)
}
