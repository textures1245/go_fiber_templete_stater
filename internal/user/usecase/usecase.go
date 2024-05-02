package usecase

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/textures1245/go-template/internal/user"
	"github.com/textures1245/go-template/internal/user/dtos"
	"github.com/textures1245/go-template/internal/user/entities"
)

type userUsecase struct {
	userRepo user.UserRepository
}

func NewUserUsecase(userRepo user.UserRepository) user.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) OnFindUser(req *entities.UserLogin) (*dtos.User, error) {

	user, err := u.userRepo.FindUser(req)
	log.Info(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) OnFetchUser() ([]dtos.User, error) {

	users, err := u.userRepo.FetchUser()
	log.Info(users)

	if err != nil {
		log.Debug("Logged 3")
		return nil, err
	}

	return users, nil
}
