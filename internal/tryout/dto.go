package tryout

import (
	"github.com/doodpanda/tryout-backend/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type TryoutListRequest struct {
	Category   string `json:"category"`
	Difficulty string `json:"difficulty"`
	Search     string `json:"search"`
}

func TryoutListRequestToParam(to TryoutListRequest, param *repository.GetTryoutListFilteredParams) error {
	param.Category = pgtype.Text{String: to.Category, Valid: false}
	param.Difficulty = pgtype.Text{String: to.Difficulty, Valid: false}
	param.Column4 = pgtype.Text{String: to.Search, Valid: false}
	return nil
}
