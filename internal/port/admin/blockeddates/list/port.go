package list

import "github.com/deday-pool-villa/backend/internal/domain"

// Repository is the port blocked-dates listing depends on.
type Repository interface {
	FindAllOrdered() ([]domain.BlockedDate, error)
}
