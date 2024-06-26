package usecase

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2/log"
	"github.com/textures1245/go-template/internal/auth"
	"github.com/textures1245/go-template/internal/auth/dtos"
	_authEntities "github.com/textures1245/go-template/internal/auth/entities"
	"github.com/textures1245/go-template/internal/user"
	_userEntities "github.com/textures1245/go-template/internal/user/entities"
	"github.com/textures1245/go-template/pkg/apperror"
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

func (u *authUse) Login(ctx context.Context, req *_authEntities.UsersCredentials) (*dtos.UserTokenRes, int, *apperror.CErr) {

	user, err := u.UsersRepo.FindUserAsPassport(ctx, req.Username)
	if err != nil {
		status, cErr := apperror.HandleAuthError(err)
		return nil, status, cErr
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		fmt.Println(err.Error())
		return nil, http.StatusBadRequest, apperror.NewCErr(apperror.ErrorInvalidCredentials, nil)
	}

	userToken, err := u.AuthRepo.SignUsersAccessToken(&_authEntities.UserSignToken{
		Id:       user.Id,
		Username: req.Username,
	})
	if err != nil {
		status, cErr := apperror.HandleAuthError(err)
		return nil, status, cErr
	}

	return userToken, http.StatusOK, nil
}

func (u *authUse) Register(ctx context.Context, req *_userEntities.UserCreatedReq) (*dtos.UsersRegisteredRes, int, *apperror.CErr) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, http.StatusInternalServerError, apperror.NewCErr(errors.New("password can not be hashed"), nil)
	}

	cred := _authEntities.UsersCredentials{
		Username: req.Username,
		Password: string(hashedPassword),
	}
	req.Password = cred.Password
	log.Info("req", req)
	user, err := u.UsersRepo.CreateUser(ctx, req)
	if err != nil {
		status, cErr := apperror.HandleAuthError(err)
		return nil, status, cErr
	}
	log.Info("res", user)

	userToken, err := u.AuthRepo.SignUsersAccessToken(&_authEntities.UserSignToken{
		Id:       *user,
		Username: req.Username,
	})
	if err != nil {
		status, cErr := apperror.HandleAuthError(err)
		return nil, status, cErr
	}

	res := &dtos.UsersRegisteredRes{
		AccessToken: userToken.AccessToken,
		CreatedAt:   userToken.IssuedAt,
		ExpiredAt:   userToken.ExpiresIn,
	}

	return res, http.StatusOK, nil

}
