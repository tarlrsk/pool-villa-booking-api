package list

import (
	"github.com/deday-pool-villa/backend/internal/domain"
	portlist "github.com/deday-pool-villa/backend/internal/port/admin/blockeddates/list"
)

type service struct{ repo portlist.Repository }

// New builds the blocked-dates listing service.
func New(repo portlist.Repository) Service { return &service{repo: repo} }

func (s *service) Execute() ([]domain.BlockedDate, error) { return s.repo.FindAllOrdered() }
