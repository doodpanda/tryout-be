package question

type QuestionSingleResponse struct {
	ID      string      `json:"id,omitempty"`
	Text    string      `json:"text,omitempty"`
	Options interface{} `json:"options,omitempty"`
	Type    string      `json:"type,omitempty"`
}

type QuestionPluralResponse struct {
	Questions []QuestionSingleResponse `json:"questions"`
}
