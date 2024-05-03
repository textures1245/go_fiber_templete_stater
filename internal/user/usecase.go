package user

import (
	"context" // Import the missing package "context"

	"github.com/textures1245/go-template/internal/user/dtos"
	"github.com/textures1245/go-template/internal/user/entities"
)

type UserUsecase interface {
	OnUserLogin(ctx context.Context, req *entities.UserLoginReq) (*dtos.UserLoginResponse, int, error) // Replace "content.Context" with "context.Context"
	OnFetchUsers(ctx context.Context) ([]*dtos.UserDetailRespond, int, error)                          // Replace "content.Context" with "context.Context"
	OnFetchUserById(ctx context.Context, userId int64) (*dtos.UserDetailRespond, int, error)
}
