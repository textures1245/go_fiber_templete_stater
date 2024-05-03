package auth

import "github.com/textures1245/go-template/internal/auth/dtos"

type AuthRepository interface {
	SignUsersAccessToken(req *struct {
		Id    int64
		Email string
	}) (*dtos.UserTokenRes, error)
}
