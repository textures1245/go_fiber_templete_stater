package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/gofiber/fiber/v2/log"
	"github.com/textures1245/go-template/internal/user"
	"github.com/textures1245/go-template/internal/user/dtos"
	"github.com/textures1245/go-template/internal/user/entities"
	"github.com/textures1245/go-template/pkg/datasource"
)

type userRepo struct {
	db   *sqlx.DB
	conn datasource.ConnTx
}

func NewUserRepository(db *sqlx.DB) user.UserRepository {
	return &userRepo{
		db:   db,
		conn: db,
	}
}

func (r *userRepo) FindUser(req *entities.UserLogin) (*dtos.User, error) {
	query := "SELECT * FROM User WHERE username = ? AND password = ?"
	row := r.db.QueryRow(query, req.Username, req.Password)
	log.Info(row)

	var user dtos.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Info(err)
		return &user, err
	}

	return &user, nil

}

func (r *userRepo) FetchUser() ([]dtos.User, error) {
	log.Debug("FetchUser")
	query := "SELECT * FROM User"
	log.Debug(r.db)
	rows, err := r.db.Query(query)
	if err != nil {
		log.Info(err)
		return nil, err
	}
	defer rows.Close()

	var users []dtos.User
	for rows.Next() {
		var user dtos.User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.CreateAt, &user.UpdateAt)
		if err != nil {
			log.Info(err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
