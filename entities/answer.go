package entities

type Answer struct {
	ID         int    `json:"id"`
	Content    string `json:"content"`
	QuestionId int    `json:"question_id"`
}
