package user

import (
	"github.com/textures1245/go-template/internal/user/dtos"
	"github.com/textures1245/go-template/internal/user/entities"
)

type UserUsecase interface {
	OnFindUser(req *entities.UserLogin) (*dtos.User, error)
	OnFetchUser() ([]dtos.User, error)
}
