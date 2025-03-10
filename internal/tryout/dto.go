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

type TryoutNewRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	LongDesc    string   `json:"long_description"`
	Category    string   `json:"category"`
	Duration    int      `json:"duration"`
	IsPublished bool     `json:"featured"`
	Topics      []string `json:"topics"`
	Difficulty  string   `json:"difficulty"`
}

func TryoutListRequestToParam(to TryoutListRequest, param *repository.GetTryoutListFilteredParams) error {
	if to.Category == "all" || to.Category == "" {
		param.Category = pgtype.Text{Valid: false}
	} else {
		param.Category = pgtype.Text{String: to.Category, Valid: true}
	}

	if to.Difficulty == "all" || to.Difficulty == "" {
		param.Difficulty = pgtype.Text{Valid: false}
	} else {
		param.Difficulty = pgtype.Text{String: to.Difficulty, Valid: true}
	}

	if to.Search == "" {
		param.Column4 = pgtype.Text{Valid: false}
	} else {
		param.Column4 = pgtype.Text{String: to.Search, Valid: true}
	}

	return nil
}

func TryoutNewRequestToParam(to TryoutNewRequest, param *repository.InsertTryoutParams) error {
	param.Title = to.Title
	param.Description = pgtype.Text{String: to.Description, Valid: true}
	param.LongDescription = pgtype.Text{String: to.LongDesc, Valid: true}
	param.Category = pgtype.Text{String: to.Category, Valid: true}
	param.Duration = pgtype.Int4{Int32: int32(to.Duration), Valid: true}
	param.IsPublished = to.IsPublished
	param.Difficulty = pgtype.Text{String: to.Difficulty, Valid: true}
	param.Topics = to.Topics
	return nil
}

func TryoutUpdateRequestToParam(to TryoutNewRequest, param *repository.UpdateTryoutParams) error {
	param.Title = to.Title
	param.Description = pgtype.Text{String: to.Description, Valid: true}
	param.LongDescription = pgtype.Text{String: to.LongDesc, Valid: true}
	param.Category = pgtype.Text{String: to.Category, Valid: true}
	param.Duration = pgtype.Int4{Int32: int32(to.Duration), Valid: true}
	param.IsPublished = to.IsPublished
	param.Difficulty = pgtype.Text{String: to.Difficulty, Valid: true}
	param.Topics = to.Topics
	return nil
}
