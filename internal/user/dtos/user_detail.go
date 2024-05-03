package dtos

type UserDetailRespond struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email" db:"email"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	UpdatedAt   string `json:"updated_at"`
	CreatedAt   string `json:"created_at"`
}
