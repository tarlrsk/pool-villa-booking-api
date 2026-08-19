package create

import "github.com/deday-pool-villa/backend/internal/domain"

// Repository is the port booking creation depends on to persist a new booking.
type Repository interface {
	Create(booking *domain.Booking) error
}
