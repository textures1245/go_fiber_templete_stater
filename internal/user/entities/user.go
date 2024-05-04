package entities

type User struct {
	Id          int64  `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Username    string `json:"username" db:"username"`
	Password    string `json:"password" db:"password"`
	Email       string `json:"email" db:"email"`
	IdCard      string `json:"id_card" db:"id_card"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
	CreatedAt   string `json:"created_at" db:"created_at"`
}
