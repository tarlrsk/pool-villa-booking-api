package list

import (
	"github.com/deday-pool-villa/backend/internal/domain"
	portlist "github.com/deday-pool-villa/backend/internal/port/admin/bookings/list"
)

type service struct{ repo portlist.Repository }

// New builds the admin bookings listing service.
func New(repo portlist.Repository) Service { return &service{repo: repo} }

func (s *service) Execute(status string) ([]domain.Booking, error) { return s.repo.FindAll(status) }
