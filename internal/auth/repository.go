package auth

import (
	"github.com/textures1245/go-template/internal/auth/dtos"
	"github.com/textures1245/go-template/internal/auth/entities"
)

type AuthRepository interface {
	SignUsersAccessToken(req *entities.UserSignToken) (*dtos.UserTokenRes, error)
}
