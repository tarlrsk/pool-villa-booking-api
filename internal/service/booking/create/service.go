package create

import (
	"github.com/deday-pool-villa/backend/internal/domain"
	bookingcreateport "github.com/deday-pool-villa/backend/internal/port/booking/create"
	availabilitycheck "github.com/deday-pool-villa/backend/internal/service/availability/check"
	pricingcalculate "github.com/deday-pool-villa/backend/internal/service/pricing/calculate"
	"github.com/deday-pool-villa/backend/internal/util"
)

type service struct {
	repo         bookingcreateport.Repository
	availability availabilitycheck.Service
	pricing      pricingcalculate.Service
}

// New builds the booking creation service.
func New(repo bookingcreateport.Repository, availability availabilitycheck.Service, pricing pricingcalculate.Service) Service {
	return &service{repo: repo, availability: availability, pricing: pricing}
}

func (s *service) Execute(input Input) (Output, error) {
	availability, err := s.availability.Execute(input.Checkin, input.Checkout)
	if err != nil {
		return Output{}, err
	}
	if !availability.Available {
		return Output{}, &UnavailableError{Reason: availability.Reason}
	}

	price, err := s.pricing.Execute(input.Checkin, input.Checkout)
	if err != nil {
		return Output{}, err
	}

	booking := domain.Booking{
		ID:          util.GenerateBookingID(),
		LineUserID:  input.LineUserID,
		DisplayName: input.DisplayName,
		Phone:       input.Phone,
		CheckIn:     input.Checkin,
		CheckOut:    input.Checkout,
		Guests:      input.Guests,
		TotalPrice:  price.TotalPrice,
		Status:      domain.StatusPending,
		Notes:       input.Notes,
	}

	if err := s.repo.Create(&booking); err != nil {
		return Output{}, err
	}

	return Output{Booking: booking, Price: price}, nil
}
