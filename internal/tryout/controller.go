package tryout

import (
	"github.com/doodpanda/tryout-backend/internal/common"
	"github.com/doodpanda/tryout-backend/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
)

type Controller interface {
	GetTryoutList(c *fiber.Ctx) error
	GetTryoutListFiltered(c *fiber.Ctx) error
	GetTryoutById(c *fiber.Ctx) error
	CreateNewTryout(c *fiber.Ctx) error
	UpdateTryout(c *fiber.Ctx) error
	DeleteTryout(c *fiber.Ctx) error
}

type controller struct {
	service Service
}

func getUserID(c *fiber.Ctx) pgtype.UUID {
	var userId pgtype.UUID
	if err := userId.Scan(c.Context().UserValue("userID").(string)); err != nil {
		_ = userId.Scan("00000000-0000-0000-0000-000000000000") // avoid nil pointer
	}

	return userId
}

func NewController(service Service) Controller {
	return &controller{
		service: service,
	}
}

func (ctr *controller) GetTryoutListFiltered(c *fiber.Ctx) error {
	var req TryoutListRequest
	var param repository.GetTryoutListFilteredParams
	userId := getUserID(c)

	param.Column1 = userId
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if err := TryoutListRequestToParam(req, &param); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	tryouts, err := ctr.service.GetTryoutListFiltered(c.Context(), param)
	if err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	TryoutList := make([]TryoutListResponse, 0)
	for _, tryout := range tryouts {
		TryoutList = append(TryoutList, TryoutListResponse{
			ID:               tryout.ID.String(),
			Title:            tryout.Title,
			Description:      tryout.Description.String,
			LongDescription:  tryout.LongDescription.String,
			Category:         tryout.Category.String,
			QuestionCount:    0,
			Duration:         int(tryout.Duration.Int32),
			CreatedAt:        tryout.CreatedAt.Time.String(),
			ParticipantCount: 0,
			Difficulty:       tryout.Difficulty.String,
			PassingScore:     int(tryout.PassingScore.Int32),
			Topics:           []string{},
			CreatorID:        tryout.CreatorID.String(),
			Featured:         tryout.IsPublished,
		})
	}

	return c.JSON(common.GeneralSuccessHasDataResponse{
		OK:      true,
		Message: "fetch tryout list success",
		Data: TryoutListPlural{
			Tryouts: TryoutList,
		},
	})
}

func (ctr *controller) GetTryoutList(c *fiber.Ctx) error {
	var uuid pgtype.UUID

	tryouts, err := ctr.service.GetTryoutList(c.Context(), uuid)
	if err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	TryoutList := make([]TryoutListResponse, 0)
	for _, tryout := range tryouts {
		TryoutList = append(TryoutList, TryoutListResponse{
			ID:               tryout.ID.String(),
			Title:            tryout.Title,
			Description:      tryout.Description.String,
			LongDescription:  tryout.LongDescription.String,
			Category:         tryout.Category.String,
			QuestionCount:    0,
			Duration:         int(tryout.Duration.Int32),
			CreatedAt:        tryout.CreatedAt.Time.String(),
			ParticipantCount: 0,
			Difficulty:       tryout.Difficulty.String,
			PassingScore:     int(tryout.PassingScore.Int32),
			Topics:           tryout.Topics,
			CreatorID:        tryout.CreatorID.String(),
			Featured:         tryout.IsPublished,
		})
	}

	return c.JSON(common.GeneralSuccessHasDataResponse{
		OK:      true,
		Message: "fetch tryout list success",
		Data: TryoutListPlural{
			Tryouts: TryoutList,
		},
	})
}

func (ctr *controller) GetTryoutById(c *fiber.Ctx) error {
	var response TryoutListResponse
	var uuid pgtype.UUID

	if err := uuid.Scan(c.Params("id")); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	tryout, err := ctr.service.GetTryoutById(c.Context(), uuid)
	if err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if err := TryoutResponse(&response, tryout); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	return c.JSON(common.GeneralSuccessHasDataResponse{
		OK:      true,
		Message: "fetch tryout success",
		Data:    TryoutListSingle{Tryout: response},
	})
}

func (ctr *controller) CreateNewTryout(c *fiber.Ctx) error {
	var req TryoutNewRequest
	var param repository.InsertTryoutParams
	userId := getUserID(c)
	if userId.String() == "00000000-0000-0000-0000-000000000000" {
		return c.Status(fiber.StatusUnauthorized).JSON(common.CreateErrorResponse(fiber.ErrUnauthorized))
	}

	param.CreatorID = getUserID(c)

	if err := c.BodyParser(&req); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if err := TryoutNewRequestToParam(req, &param); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if err := ctr.service.CreateTryout(c.Context(), param); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	return c.JSON(common.GeneralSuccessResponse{
		OK:      true,
		Message: "create new tryout success",
	})
}

func (ctr *controller) UpdateTryout(c *fiber.Ctx) error {
	var req TryoutNewRequest
	var param repository.UpdateTryoutParams

	if err := param.ID.Scan(c.Params("id")); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if err := c.BodyParser(&req); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	userId, err := ctr.service.GetTryoutCreator(c.Context(), param.ID)
	if err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if userId.String() != getUserID(c).String() {
		return c.Status(fiber.StatusUnauthorized).JSON(common.CreateErrorResponse(fiber.ErrUnauthorized))
	}

	if err := TryoutUpdateRequestToParam(req, &param); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if err := ctr.service.UpdateTryout(c.Context(), param); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	return c.JSON(common.GeneralSuccessResponse{
		OK:      true,
		Message: "update tryout success",
	})
}

func (ctr *controller) DeleteTryout(c *fiber.Ctx) error {
	uuid := pgtype.UUID{}
	if err := uuid.Scan(c.Params("id")); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	userId, err := ctr.service.GetTryoutCreator(c.Context(), uuid)
	if err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if userId.String() != getUserID(c).String() {
		return c.Status(fiber.StatusUnauthorized).JSON(common.CreateErrorResponse(fiber.ErrUnauthorized))
	}

	if err := ctr.service.DeleteTryout(c.Context(), uuid); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	return c.JSON(common.GeneralSuccessResponse{
		OK:      true,
		Message: "delete tryout success",
	})
}
