package question

import (
	"github.com/doodpanda/tryout-backend/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type QuestionCreateUpdateRequest struct {
	ID       string   `json:"id,omitempty"`
	TryoutID string   `json:"tryout_id"`
	Text     string   `json:"text,omitempty"`
	Correct  string   `json:"correctAnswer"`
	Options  []string `json:"options,omitempty"`
	Type     string   `json:"type,omitempty"`
	Points   int      `json:"points,omitempty"`
}

type MCQOptionCreate struct {
	QuestionID string `json:"question_id"`
	Option     string `json:"option"`
}

func QuestionCreateToParam(req QuestionCreateUpdateRequest, param *repository.InsertMCQQuestionParams) error {
	param.Question = req.Text
	param.Points = int32(req.Points)
	return nil
}

func QuestionUpdateToParam(req QuestionSingleResponse, param *repository.UpdateMCQQuestionParams) error {
	var questionUUID pgtype.UUID
	var correctAnswerUUID pgtype.UUID
	if err := questionUUID.Scan(req.ID); err != nil {
		return err
	}

	if err := correctAnswerUUID.Scan(req.Correct); err != nil {
		return err
	}
	param.ID = questionUUID
	param.Question = req.Text
	param.Points = int32(req.Points)
	param.CorrectAnswer = correctAnswerUUID
	return nil
}

func QuestionCreateEssayToParam(req QuestionSingleResponse, param *repository.InsertEssayQuestionParams) error {
	var tryoutUUID pgtype.UUID
	tryoutUUID.Scan(req.TryoutID)

	param.TryoutID = tryoutUUID
	param.Question = req.Text
	param.Points = int32(req.Points)
	return nil
}

func QuestionUpdateEssayToParam(req QuestionSingleResponse, param *repository.UpdateEssayQuestionParams) error {
	var questionUUID pgtype.UUID
	questionUUID.Scan(req.ID)

	param.ID = questionUUID
	param.Question = req.Text
	param.Points = int32(req.Points)
	return nil
}
