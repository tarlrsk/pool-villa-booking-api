package getunavailabledates

import "github.com/deday-pool-villa/backend/internal/domain"

// Repository is the port unavailable-date lookup depends on.
type Repository interface {
	FindBlockedDatesInRange(from, to string) ([]domain.BlockedDate, error)
	FindOverlappingBookings(from, to string) ([]domain.Booking, error)
}
