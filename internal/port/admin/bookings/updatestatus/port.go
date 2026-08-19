package updatestatus

import "github.com/deday-pool-villa/backend/internal/domain"

// Repository is the port booking status updates depends on.
type Repository interface {
	UpdateStatus(bookingID string, status domain.BookingStatus) (bool, error)
}
