package tryout

import (
	"github.com/doodpanda/tryout-backend/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func App(app fiber.Router, db *pgxpool.Pool) {
	r := app.Group("/tryout")
	var (
		repo    = repository.New(db)
		service = NewService(repo)
	)
	ctr := NewController(service)

	// Route listing

	r.Post("/", ctr.GetTryoutListFiltered)
	r.Get("/", ctr.GetTryoutList)
	r.Get("/:id", ctr.GetTryoutById)
	r.Post("/new", ctr.CreateNewTryout)
	r.Put("/:id", ctr.UpdateTryout)
	r.Delete("/:id", ctr.DeleteTryout)
}
