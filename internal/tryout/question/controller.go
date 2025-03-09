package question

import (
	"encoding/json"

	"github.com/doodpanda/tryout-backend/internal/common"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
)

type Controller interface {
	GetQuestionsByTryoutID(c *fiber.Ctx) error
}

type controller struct {
	service Service
}

func NewController(service Service) Controller {
	return &controller{
		service: service,
	}
}

func (ctr *controller) GetQuestionsByTryoutID(c *fiber.Ctx) error {
	var response QuestionPluralResponse
	var tryoutID pgtype.UUID
	if err := tryoutID.Scan(c.Params("id")); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	questions, err := ctr.service.GetQuestionsByTryoutID(c.Context(), tryoutID)
	if err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if err := json.Unmarshal(questions[0], &response.Questions); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	return c.JSON(common.GeneralSuccessHasDataResponse{
		OK:      true,
		Message: "fetch questions success",
		Data:    response,
	})
}
