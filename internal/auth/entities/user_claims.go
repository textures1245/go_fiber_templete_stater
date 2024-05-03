package entities

import (
	"github.com/golang-jwt/jwt/v4"
)

type UsersClaims struct {
	Id    int64  `json:"user_id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
