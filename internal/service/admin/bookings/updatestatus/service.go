package updatestatus

import (
	"github.com/deday-pool-villa/backend/internal/domain"
	portupdatestatus "github.com/deday-pool-villa/backend/internal/port/admin/bookings/updatestatus"
)

var validStatuses = map[domain.BookingStatus]bool{
	domain.StatusPending:   true,
	domain.StatusConfirmed: true,
	domain.StatusCancelled: true,
}

type service struct{ repo portupdatestatus.Repository }

// New builds the booking status update service.
func New(repo portupdatestatus.Repository) Service { return &service{repo: repo} }

func (s *service) Execute(bookingID string, status domain.BookingStatus) error {
	if !validStatuses[status] {
		return ErrInvalidStatus
	}
	found, err := s.repo.UpdateStatus(bookingID, status)
	if err != nil {
		return err
	}
	if !found {
		return ErrNotFound
	}
	return nil
}
