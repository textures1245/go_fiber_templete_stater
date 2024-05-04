package entities

import (
	"github.com/golang-jwt/jwt/v4"
)

type UsersClaims struct {
	Id       int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
