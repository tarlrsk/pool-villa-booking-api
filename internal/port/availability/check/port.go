package check

import "github.com/deday-pool-villa/backend/internal/domain"

// Repository is the port availability checking depends on.
type Repository interface {
	FindBlockedDatesInRange(from, to string) ([]domain.BlockedDate, error)
	CountOverlappingBookings(checkin, checkout string) (int64, error)
}
