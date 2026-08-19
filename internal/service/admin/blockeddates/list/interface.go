package list

import "github.com/deday-pool-villa/backend/internal/domain"

// Service lists all blocked dates ordered by date.
type Service interface {
	Execute() ([]domain.BlockedDate, error)
}
