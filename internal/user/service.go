package user

import (
	"context"

	"github.com/doodpanda/tryout-backend/internal/repository"
)

type Service interface {
	RegisterUser(ctx context.Context, param repository.InsertUserParams) error
	LoginUser(ctx context.Context, req string) (*repository.LoginUserRow, error)
}

type service struct {
	repo repository.Querier
}

func NewService(repo repository.Querier) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) RegisterUser(ctx context.Context, param repository.InsertUserParams) error {
	return s.repo.InsertUser(ctx, &param)
}

func (s *service) LoginUser(ctx context.Context, req string) (*repository.LoginUserRow, error) {
	return s.repo.LoginUser(ctx, req)
}
