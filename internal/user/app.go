package user

import (
	"github.com/doodpanda/tryout-backend/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func App(app fiber.Router, db *pgxpool.Pool) {
	r := app.Group("/")
	var (
		repo    = repository.New(db)
		service = NewService(repo)
	)
	ctr := NewController(service)

	// Route listing
	r.Post("/register", ctr.RegisterUser)
	r.Post("/login", ctr.LoginUser)
}
