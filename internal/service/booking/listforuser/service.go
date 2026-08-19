package listforuser

import (
	"github.com/deday-pool-villa/backend/internal/domain"
	portlistforuser "github.com/deday-pool-villa/backend/internal/port/booking/listforuser"
)

type service struct {
	repo portlistforuser.Repository
}

// New builds the user-bookings listing service.
func New(repo portlistforuser.Repository) Service {
	return &service{repo: repo}
}

func (s *service) Execute(lineUserID string) ([]domain.Booking, error) {
	return s.repo.FindByLineUserID(lineUserID)
}
