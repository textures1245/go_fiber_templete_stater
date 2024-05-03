package usecase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2/log"
	"github.com/textures1245/go-template/internal/auth"
	"github.com/textures1245/go-template/internal/auth/dtos"
	_authEntities "github.com/textures1245/go-template/internal/auth/entities"
	"github.com/textures1245/go-template/internal/user"
	_userEntities "github.com/textures1245/go-template/internal/user/entities"
	"golang.org/x/crypto/bcrypt"
)

type authUse struct {
	AuthRepo  auth.AuthRepository
	UsersRepo user.UserRepository
}

func NewAuthService(authRepo auth.AuthRepository, usersRepo user.UserRepository) auth.AuthUsecase {
	return &authUse{
		AuthRepo:  authRepo,
		UsersRepo: usersRepo,
	}
}

func (u *authUse) Login(ctx context.Context, req *_authEntities.UsersCredentials) (*dtos.UserTokenRes, int, error) {

	user, err := u.UsersRepo.FindUserAsPassport(ctx, req.Email)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		fmt.Println(err.Error())
		return nil, http.StatusBadRequest, errors.New("error, password is invalid")
	}

	userId, err := strconv.ParseInt(user.Id, 10, 64)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	userToken, err := u.AuthRepo.SignUsersAccessToken(&struct {
		Id    int64
		Email string
	}{
		Id:    userId,
		Email: req.Email,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	// res := &dtos.UsersLoginRes{
	// 	AccessToken: userToken.AccessToken,
	// 	CreatedAt:   userToken.IssuedAt,
	// 	ExpiredAt:   userToken.ExpiresIn,
	// }
	return userToken, http.StatusOK, nil
}

func (u *authUse) Register(ctx context.Context, req *_userEntities.UserCreatedReq) (*dtos.UsersRegisteredRes, int, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	cred := _authEntities.UsersCredentials{
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	req.Password = cred.Password
	log.Info("req", req)
	user, err := u.UsersRepo.CreateUser(ctx, req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	log.Info("res", user)

	userToken, err := u.AuthRepo.SignUsersAccessToken(&struct {
		Id    int64
		Email string
	}{
		Id:    *user,
		Email: req.Email,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	res := &dtos.UsersRegisteredRes{
		AccessToken: userToken.AccessToken,
		CreatedAt:   userToken.IssuedAt,
		ExpiredAt:   userToken.ExpiresIn,
	}

	return res, http.StatusOK, nil

}
