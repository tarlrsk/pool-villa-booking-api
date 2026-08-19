package login

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/deday-pool-villa/backend/internal/config"
	portlogin "github.com/deday-pool-villa/backend/internal/port/adminauth/login"
)

type service struct {
	repo portlogin.Repository
	cfg  config.Config
}

// New builds the admin login service.
func New(repo portlogin.Repository, cfg config.Config) Service {
	return &service{repo: repo, cfg: cfg}
}

// Execute verifies username/password against the Admin table and issues a JWT.
func (s *service) Execute(username, password string) (string, error) {
	admin, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", ErrInvalidCredentials
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(password)); err != nil {
		return "", ErrInvalidCredentials
	}

	expiry := time.Duration(s.cfg.AdminJWTExpiryHours) * time.Hour
	claims := Claims{
		Username: admin.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.AdminJWTSecret))
}
