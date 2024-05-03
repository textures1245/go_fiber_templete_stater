package user

import (
	"context"

	"github.com/textures1245/go-template/internal/user/entities"
)

type UserRepository interface {
	FindUserByUsernameAndPassword(ctx context.Context, req *entities.UserLoginReq) (*entities.User, error)
	CreateUser(ctx context.Context, user *entities.User) (*entities.User, error)
	GetUsers(ctx context.Context) ([]*entities.User, error)
	GetUserById(ctx context.Context, userID int64) (*entities.User, error)
}
