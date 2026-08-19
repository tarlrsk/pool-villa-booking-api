package parsetoken

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"

	"github.com/deday-pool-villa/backend/internal/config"
	"github.com/deday-pool-villa/backend/internal/service/adminauth/login"
)

type service struct {
	cfg config.Config
}

// New builds the admin token parsing service.
func New(cfg config.Config) Service {
	return &service{cfg: cfg}
}

// Execute validates a JWT and returns its claims.
func (s *service) Execute(tokenString string) (*login.Claims, error) {
	claims := &login.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.AdminJWTSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}
	return claims, nil
}
