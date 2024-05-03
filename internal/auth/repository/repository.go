package repository

import (
	"time"

	"github.com/spf13/viper"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/textures1245/go-template/internal/auth"
	"github.com/textures1245/go-template/internal/auth/dtos"
	"github.com/textures1245/go-template/internal/auth/entities"
)

type authRepo struct {
	Db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) auth.AuthRepository {
	return &authRepo{
		Db: db,
	}
}

func (r *authRepo) SignUsersAccessToken(req *struct {
	Id    int64
	Email string
}) (*dtos.UserTokenRes, error) {
	claims := entities.UsersClaims{
		Id:    req.Id,
		Email: req.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "access_token",
			Subject:   "users_access_token",
			ID:        uuid.NewString(),
			Audience:  []string{"users"},
		},
	}

	mySigningKey := viper.GetString("JWT_SECRET_TOKEN")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(mySigningKey))
	if err != nil {
		return nil, err
	}
	return &dtos.UserTokenRes{
		AccessToken: ss,
		TokenType:   "Authorization",
		ExpiresIn:   claims.ExpiresAt.String(),
		IssuedAt:    claims.IssuedAt.String(),
	}, nil
}
