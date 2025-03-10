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

func NewController(service Service) Controller {
	return &controller{
		service: service,
	}
}

func (ctr *controller) GetTryoutListFiltered(c *fiber.Ctx) error {
	var req TryoutListRequest
	var param repository.GetTryoutListFilteredParams
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
	uuid := pgtype.UUID{}
	uuid.Scan("aa645725-37e1-41ed-a7a9-71f428b05b1a")
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
	uuid := pgtype.UUID{}
	uuid.Scan(c.Params("id"))
	tryout, err := ctr.service.GetTryoutById(c.Context(), uuid)

	if err := TryoutResponse(&response, tryout); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if err != nil {
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
	param.CreatorID.Scan("aa645725-37e1-41ed-a7a9-71f428b05b1a")

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
	if err := ctr.service.DeleteTryout(c.Context(), uuid); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	return c.JSON(common.GeneralSuccessResponse{
		OK:      true,
		Message: "delete tryout success",
	})
}
