package entities

type Question struct {
	ID      int      `json:"id"`
	Content string   `json:"content"`
	Options []string `json:"options"`
}
