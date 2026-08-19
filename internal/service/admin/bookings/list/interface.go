package list

import "github.com/deday-pool-villa/backend/internal/domain"

// Service lists all bookings, optionally filtered by status.
type Service interface {
	Execute(status string) ([]domain.Booking, error)
}
