package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/gofiber/fiber/v2/log"
	_authEntities "github.com/textures1245/go-template/internal/auth/entities"
	"github.com/textures1245/go-template/internal/user"
	"github.com/textures1245/go-template/internal/user/entities"
	"github.com/textures1245/go-template/internal/user/repository/repository_query"
	"github.com/textures1245/go-template/pkg/utils"
	// "github.com/textures1245/go-template/pkg/datasource"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) user.UserRepository {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) FindUserAsPassport(ctx context.Context, username string) (*_authEntities.UsersPassport, error) {
	// checking if user email was founded

	userData := &entities.User{}

	err := r.db.QueryRowx(repository_query.FindUserByUsername, username).StructScan(userData)
	log.Info(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("UserUniqueUsernameNotFound")
		} else {
			log.Error(err)
			return nil, err

		}
	}

	userPassport := &_authEntities.UsersPassport{
		Id:       userData.Id,
		Username: userData.Username,
		Password: userData.Password,
	}

	return userPassport, nil
}

func (r *userRepo) FindUserByUsernameAndPassword(ctx context.Context, req *entities.UserLoginReq) (*entities.User, error) {
	userData := &entities.User{}
	err := r.db.QueryRowxContext(ctx, repository_query.FindUserById, req).StructScan(userData)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return userData, nil

}

func (r *userRepo) GetUsers(ctx context.Context) ([]*entities.User, error) {
	users := []*entities.User{}

	rows, err := r.db.QueryxContext(ctx, repository_query.GetUsers)
	if err != nil {
		log.Info(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entities.User
		err := rows.StructScan(&user)
		if err != nil {
			log.Info(err)
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *userRepo) GetUserById(ctx context.Context, userID int64) (*entities.User, error) {
	userData := &entities.User{}

	// query := "SELECT * FROM User WHERE id = ?"
	err := r.db.QueryRowxContext(ctx, repository_query.FindUserById, userID).StructScan(userData)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return userData, nil
}

func (r *userRepo) CreateUser(ctx context.Context, user *entities.UserCreatedReq) (*int64, error) {

	args := utils.Array{
		user.Name,
		user.Username,
		user.Password,
		user.Email,
		user.PhoneNumber,
		user.IdCard,
	}

	log.Info(args)

	res, err := r.db.ExecContext(ctx, repository_query.InsertUser, args...)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	userId, _ := res.RowsAffected()

	return &userId, nil
}

func (r *userRepo) UpdateUserById(ctx context.Context, userID int64, user *entities.UserUpdateReq) error {
	args := utils.Array{
		user.Name,
		user.Email,
		user.IdCard,
		user.PhoneNumber,
		userID,
	}

	log.Info(args)

	res, err := r.db.ExecContext(ctx, repository_query.UpdateUserById, args...)
	if err != nil {
		log.Error(err)
		return err
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *userRepo) DeleteUserById(ctx context.Context, userID int64) error {
	res, err := r.db.ExecContext(ctx, repository_query.DeleteUserById, userID)
	if err != nil {
		log.Error(err)
		return err
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// TODO: Checking row affect for the update and delete
