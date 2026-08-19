package list

import "github.com/deday-pool-villa/backend/internal/domain"

// Service lists all custom periods ordered by start date.
type Service interface {
	Execute() ([]domain.CustomPeriod, error)
}
