package listforuser

import "github.com/deday-pool-villa/backend/internal/domain"

// Service lists all bookings for a given LINE user.
type Service interface {
	Execute(lineUserID string) ([]domain.Booking, error)
}
