package entities

type User struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Username    string `json:"username" db:"username"`
	Password    string `json:"password" db:"password"`
	Email       string `json:"email" db:"email"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
	CreatedAt   string `json:"created_at" db:"created_at"`
}
