package user

import (
	"context"

	_authEntities "github.com/textures1245/go-template/internal/auth/entities"
	_userEntities "github.com/textures1245/go-template/internal/user/entities"
)

type UserRepository interface {
	FindUserByUsernameAndPassword(ctx context.Context, req *_userEntities.UserLoginReq) (*_userEntities.User, error)
	CreateUser(ctx context.Context, user *_userEntities.UserCreatedReq) (*int64, error) // Fixed the mixed named and unnamed parameters
	GetUsers(ctx context.Context) ([]*_userEntities.User, error)
	FindUserAsPassport(ctx context.Context, email string) (*_authEntities.UsersPassport, error)
	GetUserById(ctx context.Context, userID int64) (*_userEntities.User, error)
}
