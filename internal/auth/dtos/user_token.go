package dtos

type UserTokenRes struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
	IssuedAt    string `json:"issued_at"`
}
