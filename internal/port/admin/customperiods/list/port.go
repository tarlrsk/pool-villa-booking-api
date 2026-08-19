package list

import "github.com/deday-pool-villa/backend/internal/domain"

// Repository is the port custom-periods listing depends on.
type Repository interface {
	FindAllOrdered() ([]domain.CustomPeriod, error)
}
