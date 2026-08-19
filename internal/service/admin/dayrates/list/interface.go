package list

import "github.com/deday-pool-villa/backend/internal/domain"

// Service lists all day rates.
type Service interface {
	Execute() ([]domain.DayRate, error)
}
