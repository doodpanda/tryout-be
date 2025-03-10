package question

import (
	"context"

	"github.com/doodpanda/tryout-backend/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service interface {
	GetQuestionsByTryoutID(ctx context.Context, tryoutID pgtype.UUID) ([][]byte, error)
	GetQuestionByID(ctx context.Context, questionID pgtype.UUID) ([]byte, error)
	CreateMCQQuestion(ctx context.Context, param repository.InsertMCQQuestionParams) (pgtype.UUID, error)
	UpdateMCQQuestion(ctx context.Context, param repository.UpdateMCQQuestionParams) error
	CreateEssayQuestion(ctx context.Context, param repository.InsertEssayQuestionParams) error
	UpdateEssayQuestion(ctx context.Context, param repository.UpdateEssayQuestionParams) error
	DeleteQuestion(ctx context.Context, questionID pgtype.UUID) error
	CreateMCQOption(ctx context.Context, param repository.InsertOptionParams) (pgtype.UUID, error)
	GetTryoutCreator(ctx context.Context, tryoutID pgtype.UUID) (pgtype.UUID, error)
	GetTryoutCreatorByQuestionID(ctx context.Context, questionID pgtype.UUID) (pgtype.UUID, error)
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

func (s *service) GetQuestionByID(ctx context.Context, questionID pgtype.UUID) ([]byte, error) {
	return s.repo.GetQuestionByID(ctx, questionID)
}

func (s *service) CreateMCQQuestion(ctx context.Context, param repository.InsertMCQQuestionParams) (pgtype.UUID, error) {
	return s.repo.InsertMCQQuestion(ctx, &param)
}

func (s *service) UpdateMCQQuestion(ctx context.Context, param repository.UpdateMCQQuestionParams) error {
	return s.repo.UpdateMCQQuestion(ctx, &param)
}

func (s *service) CreateEssayQuestion(ctx context.Context, param repository.InsertEssayQuestionParams) error {
	return s.repo.InsertEssayQuestion(ctx, &param)
}

func (s *service) UpdateEssayQuestion(ctx context.Context, param repository.UpdateEssayQuestionParams) error {
	return s.repo.UpdateEssayQuestion(ctx, &param)
}

func (s *service) DeleteQuestion(ctx context.Context, questionID pgtype.UUID) error {
	k := s.repo.DeleteEssayQuestion(ctx, questionID)
	if k != nil {
		return k
	}
	return s.repo.DeleteQuestion(ctx, questionID)
}

func (s *service) CreateMCQOption(ctx context.Context, param repository.InsertOptionParams) (pgtype.UUID, error) {
	return s.repo.InsertOption(ctx, &param)
}

func (s *service) GetTryoutCreator(ctx context.Context, tryoutID pgtype.UUID) (pgtype.UUID, error) {
	return s.repo.GetTryoutCreator(ctx, tryoutID)
}

func (s *service) GetTryoutCreatorByQuestionID(ctx context.Context, questionID pgtype.UUID) (pgtype.UUID, error) {
	return s.repo.GetTryoutCreatorByQuestionID(ctx, questionID)
}
