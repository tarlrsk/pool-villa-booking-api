package list

import "github.com/deday-pool-villa/backend/internal/domain"

// Repository is the port admin bookings listing depends on.
type Repository interface {
	FindAll(status string) ([]domain.Booking, error)
}
