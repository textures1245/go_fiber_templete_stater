package entities

type UserCreatedReq struct {
	Name        string `json:"name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
