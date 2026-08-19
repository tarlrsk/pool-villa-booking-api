package list

import (
	"github.com/deday-pool-villa/backend/internal/domain"
	portlist "github.com/deday-pool-villa/backend/internal/port/admin/dayrates/list"
)

type service struct{ repo portlist.Repository }

// New builds the day-rates listing service.
func New(repo portlist.Repository) Service { return &service{repo: repo} }

func (s *service) Execute() ([]domain.DayRate, error) { return s.repo.FindAll() }
