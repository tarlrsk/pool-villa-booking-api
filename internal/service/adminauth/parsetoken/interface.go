package parsetoken

import "github.com/deday-pool-villa/backend/internal/service/adminauth/login"

// Service validates an admin JWT and returns its claims.
type Service interface {
	Execute(tokenString string) (*login.Claims, error)
}
