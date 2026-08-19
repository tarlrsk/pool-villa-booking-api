package sendbookingconfirmation

import "github.com/deday-pool-villa/backend/internal/domain"

// Service sends a LINE push message confirming a new booking.
type Service interface {
	Execute(booking domain.Booking) error
}
