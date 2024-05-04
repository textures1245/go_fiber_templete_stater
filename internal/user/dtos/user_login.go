package dtos

type UserLoginResponse struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}
