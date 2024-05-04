package dtos

type UserDetailRespond struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email" db:"email"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	IdCard      string `json:"id_card"`
	UpdatedAt   string `json:"updated_at"`
	CreatedAt   string `json:"created_at"`
}
