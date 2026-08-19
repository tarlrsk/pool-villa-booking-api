package upsert

import (
	"github.com/deday-pool-villa/backend/internal/domain"
	portupsert "github.com/deday-pool-villa/backend/internal/port/admin/dayrates/upsert"
)

type service struct{ repo portupsert.Repository }

// New builds the day-rates upsert service.
func New(repo portupsert.Repository) Service { return &service{repo: repo} }

func (s *service) Execute(rates []domain.DayRate) error {
	if len(rates) == 0 {
		return ErrEmpty
	}
	return s.repo.UpsertAll(rates)
}
