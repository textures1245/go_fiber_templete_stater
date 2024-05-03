package usecase

import (
	"context"
	"net/http"

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

func (u *userUsecase) OnUserLogin(ctx context.Context, req *entities.UserLoginReq) (*dtos.UserLoginResponse, int, error) {

	user, err := u.userRepo.FindUserByUsernameAndPassword(ctx, req)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return entities.NewUserLogin(user), http.StatusOK, nil
}

func (u *userUsecase) OnFetchUsers(ctx context.Context) ([]*dtos.UserDetailRespond, int, error) {

	users, err := u.userRepo.GetUsers(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// the problem was found here when trying to casting the user to the UserDetailRespond
	var res []*dtos.UserDetailRespond
	for _, user := range users {

		res = append(res, entities.NewUserDetail(user))
	}

	return res, http.StatusOK, nil
}

func (u *userUsecase) OnFetchUserById(ctx context.Context, userId int64) (*dtos.UserDetailRespond, int, error) {

	user, err := u.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return entities.NewUserDetail(user), http.StatusOK, nil
}
