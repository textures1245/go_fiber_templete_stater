package auth

import (
	"context"

	"github.com/textures1245/go-template/internal/auth/dtos"
	_authEntities "github.com/textures1245/go-template/internal/auth/entities"
	_userEntities "github.com/textures1245/go-template/internal/user/entities"
)

type AuthUsecase interface {
	Login(ctx context.Context, req *_authEntities.UsersCredentials) (*dtos.UserTokenRes, int, error)
	Register(ctx context.Context, req *_userEntities.UserCreatedReq) (*dtos.UsersRegisteredRes, int, error)
}
