package create

import "github.com/deday-pool-villa/backend/internal/domain"

// Repository is the port blocked-date creation depends on.
type Repository interface {
	Create(blocked *domain.BlockedDate) error
}
