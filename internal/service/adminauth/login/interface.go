package login

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var ErrInvalidCredentials = errors.New("invalid username or password")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Service authenticates an admin and issues a JWT.
type Service interface {
	Execute(username, password string) (string, error)
}
