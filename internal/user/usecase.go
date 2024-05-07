package user

import (
	"context" // Import the missing package "context"

	"github.com/textures1245/go-template/internal/user/dtos"
	"github.com/textures1245/go-template/internal/user/entities"
	"github.com/textures1245/go-template/pkg/apperror"
)

type UserUsecase interface {
	OnUserLogin(ctx context.Context, req *entities.UserLoginReq) (*dtos.UserLoginResponse, int, *apperror.CErr) // Replace "content.Context" with "context.Context"
	OnFetchUsers(ctx context.Context) ([]*dtos.UserDetailRespond, int, *apperror.CErr)                          // Replace "content.Context" with "context.Context"
	OnFetchUserById(ctx context.Context, userId int64) (*dtos.UserDetailRespond, int, *apperror.CErr)
	OnUpdateUserById(ctx context.Context, userId int64, req *entities.UserUpdateReq) (int, *apperror.CErr)
	UserDeleted(ctx context.Context, userId int64) (int, *apperror.CErr)
}
