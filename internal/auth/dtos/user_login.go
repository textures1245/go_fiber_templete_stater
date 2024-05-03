package dtos

type UsersRegisteredRes struct {
	AccessToken string `json:"access_token"`
	CreatedAt   string `json:"created_at"`
	ExpiredAt   string `json:"expired_at"`
}
