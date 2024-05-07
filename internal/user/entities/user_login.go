package entities

import "github.com/textures1245/go-template/internal/user/dtos"

type UserLoginReq struct {
	Username string `json:"username" form:"username" validate:"required,min=5,max=50" binding:"required"`
	Password string `json:"password" form:"password"  validate:"required,min=8" binding:"required"`
}

func NewUserLogin(dat *User) *dtos.UserLoginResponse {
	return &dtos.UserLoginResponse{
		Id:        dat.Id,
		Name:      dat.Name,
		Email:     dat.Email,
		UpdatedAt: dat.UpdatedAt,
		CreatedAt: dat.CreatedAt,
	}
}
