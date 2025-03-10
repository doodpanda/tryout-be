package user

import (
	"time"

	"github.com/doodpanda/tryout-backend/internal/common"
	"github.com/doodpanda/tryout-backend/internal/repository"
	"github.com/doodpanda/tryout-backend/internal/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type Controller interface {
	RegisterUser(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
}

type controller struct {
	service Service
}

func NewController(service Service) Controller {
	return &controller{
		service: service,
	}
}

func (ctr *controller) RegisterUser(c *fiber.Ctx) error {
	var req RegisterUserRequest
	var param repository.InsertUserParams
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if err := RegisterUserToParam(req, &param); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if err := ctr.service.RegisterUser(c.Context(), param); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	return c.JSON(common.GeneralSuccessResponse{
		OK:      true,
		Message: "register user success",
	})
}

func (ctr *controller) LoginUser(c *fiber.Ctx) error {
	var req LoginUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	userInfo, err := ctr.service.LoginUser(c.Context(), req.Email)
	if err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(req.Password)); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(common.CreateErrorResponse(err))
	}

	token, err := utils.GenerateToken(userInfo.ID.String())
	if err != nil {
		return c.JSON(common.CreateErrorResponse(err))
	}

	c.Cookie(&fiber.Cookie{
		Name:     "Authorization",
		Value:    "Bearer " + token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	})

	return c.JSON(common.GeneralSuccessResponse{
		OK:      true,
		Message: "login user success",
	})
}
