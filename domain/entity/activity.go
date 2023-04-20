package entity

// Activity is a struct for activity domain
type Activity struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Email     string `json:"email"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}
