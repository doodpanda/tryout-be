package question

import (
	"encoding/json"

	"github.com/doodpanda/tryout-backend/internal/common"
	"github.com/doodpanda/tryout-backend/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
)

type Controller interface {
	GetQuestionsByTryoutID(c *fiber.Ctx) error
	GetSingleQuestion(c *fiber.Ctx) error
	CreateQuestion(c *fiber.Ctx) error
	UpdateQuestion(c *fiber.Ctx) error
	DeleteQuestion(c *fiber.Ctx) error
	CreateEssayQuestion(c *fiber.Ctx) error
	UpdateEssayQuestion(c *fiber.Ctx) error
}

func getUserID(c *fiber.Ctx) pgtype.UUID {
	var userId pgtype.UUID
	if err := userId.Scan(c.Context().UserValue("userID").(string)); err != nil {
		_ = userId.Scan("00000000-0000-0000-0000-000000000000") // avoid nil pointer
	}

	return userId
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
	}

	return c.JSON(common.GeneralSuccessHasDataResponse{
		OK:      true,
		Message: "fetch questions success",
		Data:    response,
	})
}

func (ctr *controller) GetSingleQuestion(c *fiber.Ctx) error {
	var response QuestionSingleResponseToSend
	var questionID pgtype.UUID
	if err := questionID.Scan(c.Params("question_id")); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	question, err := ctr.service.GetQuestionByID(c.Context(), questionID)
	if err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}
	if err := json.Unmarshal(question, &response); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	return c.JSON(common.GeneralSuccessHasDataResponse{
		OK:      true,
		Message: "fetch question success",
		Data:    response,
	})
}

func (ctr *controller) CreateQuestion(c *fiber.Ctx) error {
	var req QuestionCreateUpdateRequest
	var param repository.InsertMCQQuestionParams
	var id pgtype.UUID

	if err := c.BodyParser(&req); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if err := QuestionCreateToParam(req, &param); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if err := param.TryoutID.Scan(c.Params("id")); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	userId, err := ctr.service.GetTryoutCreator(c.Context(), param.TryoutID)
	if err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if userId.String() != getUserID(c).String() {
		return c.Status(fiber.StatusUnauthorized).JSON(common.CreateErrorResponse(fiber.ErrUnauthorized))
	}

	for _, option := range req.Options {
		optionID, err := ctr.service.CreateMCQOption(c.Context(), repository.InsertOptionParams{
			QuestionID: id,
			Option:     option,
		})
		if err != nil {
			return c.JSON(common.CreateErrorResponse(err))
		}
		if option == req.Correct {
			err := ctr.service.UpdateMCQQuestion(c.Context(), repository.UpdateMCQQuestionParams{
				ID:            id,
				CorrectAnswer: optionID,
				Points:        int32(req.Points),
				Question:      req.Text,
			})

			if err != nil {
				return c.JSON(common.CreateErrorResponse(err))
			}

		}
	}

	return c.JSON(common.GeneralSuccessResponse{
		OK:      true,
		Message: "create question success",
	})
}

func (ctr *controller) UpdateQuestion(c *fiber.Ctx) error {
	var req QuestionSingleResponse
	var param repository.UpdateMCQQuestionParams
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if err := QuestionUpdateToParam(req, &param); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}
	userId, err := ctr.service.GetTryoutCreatorByQuestionID(c.Context(), param.ID)
	if err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if userId.String() != getUserID(c).String() {
		return c.Status(fiber.StatusUnauthorized).JSON(common.CreateErrorResponse(fiber.ErrUnauthorized))
	}

	if err := ctr.service.UpdateMCQQuestion(c.Context(), param); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	return c.JSON(common.GeneralSuccessResponse{
		OK:      true,
		Message: "update question success",
	})
}

func (ctr *controller) CreateEssayQuestion(c *fiber.Ctx) error {
	var req QuestionSingleResponse
	var param repository.InsertEssayQuestionParams

	if err := c.BodyParser(&req); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if err := QuestionCreateEssayToParam(req, &param); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	userId, err := ctr.service.GetTryoutCreator(c.Context(), param.TryoutID)
	if err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if userId.String() != getUserID(c).String() {
		return c.Status(fiber.StatusUnauthorized).JSON(common.CreateErrorResponse(fiber.ErrUnauthorized))
	}

	if err := ctr.service.CreateEssayQuestion(c.Context(), param); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	return c.JSON(common.GeneralSuccessResponse{
		OK:      true,
		Message: "create essay question success",
	})
}

func (ctr *controller) UpdateEssayQuestion(c *fiber.Ctx) error {
	var req QuestionSingleResponse
	var param repository.UpdateEssayQuestionParams
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if err := QuestionUpdateEssayToParam(req, &param); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	userId, err := ctr.service.GetTryoutCreatorByQuestionID(c.Context(), param.ID)
	if err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if userId.String() != getUserID(c).String() {
		return c.Status(fiber.StatusUnauthorized).JSON(common.CreateErrorResponse(fiber.ErrUnauthorized))
	}

	if err := ctr.service.UpdateEssayQuestion(c.Context(), param); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	return c.JSON(common.GeneralSuccessResponse{
		OK:      true,
		Message: "update essay question success",
	})
}

func (ctr *controller) DeleteQuestion(c *fiber.Ctx) error {
	var questionID pgtype.UUID
	if err := questionID.Scan(c.Params("question_id")); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	userId, err := ctr.service.GetTryoutCreatorByQuestionID(c.Context(), questionID)
	if err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if userId.String() != getUserID(c).String() {
		return c.Status(fiber.StatusUnauthorized).JSON(common.CreateErrorResponse(fiber.ErrUnauthorized))
	}

	if err := ctr.service.DeleteQuestion(c.Context(), questionID); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	return c.JSON(common.GeneralSuccessResponse{
		OK:      true,
		Message: "delete question success",
	})
}

func (ctr *controller) CreateMCQOption(c *fiber.Ctx) error {
	var req repository.InsertOptionParams
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	userId, err := ctr.service.GetTryoutCreatorByQuestionID(c.Context(), req.QuestionID)
	if err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if userId.String() != getUserID(c).String() {
		return c.Status(fiber.StatusUnauthorized).JSON(common.CreateErrorResponse(fiber.ErrUnauthorized))
	}

	id, err := ctr.service.CreateMCQOption(c.Context(), req)
	if err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	return c.JSON(common.GeneralSuccessResponse{
		OK:      true,
		Message: id.String(),
	})
}
