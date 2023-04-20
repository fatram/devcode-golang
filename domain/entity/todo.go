package entity

// Todo is a struct for todo domain
type Todo struct {
	ID         int    `json:"id"`
	ActivityID int    `json:"activity_group_id"`
	Title      string `json:"title"`
	IsActive   bool   `json:"is_active"`
	Priority   string `json:"priority"`
	CreatedAt  int    `json:"created_at"`
	UpdatedAt  int    `json:"updated_at"`
}
