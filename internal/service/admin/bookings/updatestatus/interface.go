package updatestatus

import (
	"errors"

	"github.com/deday-pool-villa/backend/internal/domain"
)

var (
	ErrInvalidStatus = errors.New("invalid status")
	ErrNotFound      = errors.New("booking not found")
)

// Service updates a booking's status.
type Service interface {
	Execute(bookingID string, status domain.BookingStatus) error
}
