package auth

import (
	"context"

	"github.com/textures1245/go-template/internal/auth/dtos"
	_authEntities "github.com/textures1245/go-template/internal/auth/entities"
	_userEntities "github.com/textures1245/go-template/internal/user/entities"
	"github.com/textures1245/go-template/pkg/apperror"
)

type AuthUsecase interface {
	Login(ctx context.Context, req *_authEntities.UsersCredentials) (*dtos.UserTokenRes, int, *apperror.CErr)
	Register(ctx context.Context, req *_userEntities.UserCreatedReq) (*dtos.UsersRegisteredRes, int, *apperror.CErr)
}
