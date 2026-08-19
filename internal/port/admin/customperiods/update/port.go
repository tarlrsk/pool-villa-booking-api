package update

import "github.com/deday-pool-villa/backend/internal/domain"

// Repository is the port custom-period update depends on.
type Repository interface {
	UpdateByID(id string, updates domain.CustomPeriod) (bool, error)
}
