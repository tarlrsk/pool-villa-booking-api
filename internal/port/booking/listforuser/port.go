package listforuser

import "github.com/deday-pool-villa/backend/internal/domain"

// Repository is the port user-bookings listing depends on.
type Repository interface {
	FindByLineUserID(lineUserID string) ([]domain.Booking, error)
}
