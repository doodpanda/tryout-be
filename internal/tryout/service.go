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
	CreateTryout(ctx context.Context, param repository.InsertTryoutParams) error
	UpdateTryout(ctx context.Context, param repository.UpdateTryoutParams) error
	DeleteTryout(ctx context.Context, tryoutID pgtype.UUID) error
	GetTryoutCreator(ctx context.Context, tryoutID pgtype.UUID) (pgtype.UUID, error)
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

func (s *service) CreateTryout(ctx context.Context, param repository.InsertTryoutParams) error {
	return s.repo.InsertTryout(ctx, &param)
}

func (s *service) UpdateTryout(ctx context.Context, param repository.UpdateTryoutParams) error {
	return s.repo.UpdateTryout(ctx, &param)
}

func (s *service) DeleteTryout(ctx context.Context, tryoutID pgtype.UUID) error {
	return s.repo.DeleteTryout(ctx, tryoutID)
}

func (s *service) GetTryoutCreator(ctx context.Context, tryoutID pgtype.UUID) (pgtype.UUID, error) {
	return s.repo.GetTryoutCreator(ctx, tryoutID)
}
