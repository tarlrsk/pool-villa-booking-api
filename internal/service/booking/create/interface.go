package create

import (
	"github.com/deday-pool-villa/backend/internal/domain"
	pricingcalculate "github.com/deday-pool-villa/backend/internal/service/pricing/calculate"
)

type Input struct {
	LineUserID  string
	DisplayName string
	Phone       string
	Checkin     string
	Checkout    string
	Guests      int
	Notes       string
}

type Output struct {
	Booking domain.Booking
	Price   pricingcalculate.Result
}

// UnavailableError is returned when the requested stay cannot be booked.
type UnavailableError struct {
	Reason string
}

func (e *UnavailableError) Error() string {
	if e.Reason == "" {
		return "dates not available"
	}
	return e.Reason
}

// Service creates a new booking after checking availability and pricing it.
type Service interface {
	Execute(input Input) (Output, error)
}
