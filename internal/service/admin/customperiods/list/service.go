package list

import (
	"github.com/deday-pool-villa/backend/internal/domain"
	portlist "github.com/deday-pool-villa/backend/internal/port/admin/customperiods/list"
)

type service struct{ repo portlist.Repository }

// New builds the custom-periods listing service.
func New(repo portlist.Repository) Service { return &service{repo: repo} }

func (s *service) Execute() ([]domain.CustomPeriod, error) { return s.repo.FindAllOrdered() }
