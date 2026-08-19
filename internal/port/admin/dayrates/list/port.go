package list

import "github.com/deday-pool-villa/backend/internal/domain"

// Repository is the port day-rates listing depends on.
type Repository interface {
	FindAll() ([]domain.DayRate, error)
}
