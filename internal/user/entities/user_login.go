package entities

import "github.com/textures1245/go-template/internal/user/dtos"

type UserLoginReq struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func NewUserLogin(dat *User) *dtos.UserLoginResponse {
	return &dtos.UserLoginResponse{
		ID:        dat.ID,
		Name:      dat.Name,
		Email:     dat.Email,
		UpdatedAt: dat.UpdatedAt,
		CreatedAt: dat.CreatedAt,
	}
}
