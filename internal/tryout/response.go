package tryout

import "github.com/doodpanda/tryout-backend/internal/repository"

type TryoutListResponse struct {
	ID               string   `json:"id"`
	Title            string   `json:"title"`
	Description      string   `json:"description"`
	LongDescription  string   `json:"long_description"`
	Category         string   `json:"category"`
	QuestionCount    int      `json:"question_count"`
	Duration         int      `json:"duration"`
	CreatedAt        string   `json:"created_at"`
	ParticipantCount int      `json:"participants"`
	Difficulty       string   `json:"difficulty"`
	PassingScore     int      `json:"passing_score"`
	Topics           []string `json:"topics"`
	CreatorID        string   `json:"creator_id"`
	Featured         bool     `json:"featured"`
}

type TryoutListPlural struct {
	Tryouts []TryoutListResponse `json:"tryouts"`
}

type TryoutListSingle struct {
	Tryout TryoutListResponse `json:"tryout"`
}

func TryoutResponse(tryout *TryoutListResponse, param *repository.Tryout) error {
	tryout.ID = param.ID.String()
	tryout.Title = param.Title
	tryout.Description = param.Description.String
	tryout.LongDescription = param.LongDescription.String
	tryout.Category = param.Category.String
	tryout.QuestionCount = 0
	tryout.Duration = int(param.Duration.Int32)
	tryout.CreatedAt = param.CreatedAt.Time.String()
	tryout.ParticipantCount = 0
	tryout.Difficulty = param.Difficulty.String
	tryout.PassingScore = int(param.PassingScore.Int32)
	tryout.Topics = param.Topics
	tryout.CreatorID = param.CreatorID.String()
	tryout.Featured = param.IsPublished
	return nil
}
