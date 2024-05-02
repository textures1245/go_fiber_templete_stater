package user

import (
	"github.com/textures1245/go-template/internal/user/dtos"
	"github.com/textures1245/go-template/internal/user/entities"
)

type UserRepository interface {
	FindUser(req *entities.UserLogin) (*dtos.User, error)
	FetchUser() ([]dtos.User, error)
}
