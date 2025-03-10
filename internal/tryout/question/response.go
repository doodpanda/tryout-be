package question

type QuestionSingleResponse struct {
	ID       string      `json:"id,omitempty"`
	TryoutID string      `json:"tryout_id"`
	Text     string      `json:"text,omitempty"`
	Correct  string      `json:"correct_answer"`
	Options  interface{} `json:"options,omitempty"`
	Type     string      `json:"type,omitempty"`
	Points   int         `json:"points,omitempty"`
}

type QuestionPluralResponse struct {
	Questions []QuestionSingleResponse `json:"questions"`
}

type QuestionSingleResponseToSend struct {
	Question QuestionSingleResponse `json:"question"`
}
