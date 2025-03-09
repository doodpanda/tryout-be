package question

import (
	"context"

	"github.com/doodpanda/tryout-backend/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service interface {
	GetQuestionsByTryoutID(ctx context.Context, tryoutID pgtype.UUID) ([][]byte, error)
}

type service struct {
	repo repository.Querier
}

func NewService(repo repository.Querier) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetQuestionsByTryoutID(ctx context.Context, tryoutID pgtype.UUID) ([][]byte, error) {
	return s.repo.GetTryoutQuestionsByTryoutId(ctx, tryoutID)
}
