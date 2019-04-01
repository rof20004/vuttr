package models

// Tool model
type Tool struct {
	ID          int64    `json:"id"`
	Title       string   `json:"title"`
	Link        string   `json:"link"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}
