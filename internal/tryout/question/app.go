package question

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
	r.Get("/:id/questions", ctr.GetQuestionsByTryoutID)
	r.Get("/:id/questions/:question_id", ctr.GetSingleQuestion)
	r.Post("/:id/questions/", ctr.CreateQuestion)
	r.Put("/:id/questions/:question_id", ctr.UpdateQuestion)
	r.Delete("/:id/questions/:question_id", ctr.DeleteQuestion)
}
