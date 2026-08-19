package login

import "github.com/deday-pool-villa/backend/internal/domain"

// Repository is the port admin login depends on to look up the account.
type Repository interface {
	FindByUsername(username string) (domain.Admin, error)
}
