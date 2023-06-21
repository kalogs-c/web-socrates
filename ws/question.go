package ws

type Question struct {
	ID            int      `json:"id"`
	Text          string   `json:"text"`
	Answers       []string `json:"answers"`
	Subject       string   `json:"subject"`
	CorrectAnswer int8     `json:"correct_answer"`
}

func (q *Question) GetRandomQuestion(subject string) *Question {
	return &Question{
		ID:   1,
		Text: "Question",
		Answers: []string{
			"Answer 1",
			"Answer 2",
			"Answer 3",
			"Answer 4",
		},
		Subject:       subject,
		CorrectAnswer: 2,
	}
}
