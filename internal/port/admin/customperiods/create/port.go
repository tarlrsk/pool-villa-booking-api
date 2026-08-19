package create

import "github.com/deday-pool-villa/backend/internal/domain"

// Repository is the port custom-period creation depends on.
type Repository interface {
	Create(period *domain.CustomPeriod) error
}
