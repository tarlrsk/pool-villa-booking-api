package create

import (
	"github.com/deday-pool-villa/backend/internal/domain"
	portcreate "github.com/deday-pool-villa/backend/internal/port/admin/blockeddates/create"
)

type service struct{ repo portcreate.Repository }

// New builds the blocked-date creation service.
func New(repo portcreate.Repository) Service { return &service{repo: repo} }

func (s *service) Execute(input Input) (domain.BlockedDate, error) {
	blocked := domain.BlockedDate{Date: input.Date, Reason: input.Reason}
	if err := s.repo.Create(&blocked); err != nil {
		return domain.BlockedDate{}, err
	}
	return blocked, nil
}
