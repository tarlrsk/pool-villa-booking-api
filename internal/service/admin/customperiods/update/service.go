package update

import (
	"github.com/deday-pool-villa/backend/internal/domain"
	portupdate "github.com/deday-pool-villa/backend/internal/port/admin/customperiods/update"
)

type service struct{ repo portupdate.Repository }

// New builds the custom-period update service.
func New(repo portupdate.Repository) Service { return &service{repo: repo} }

func (s *service) Execute(input Input) error {
	updates := domain.CustomPeriod{
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		Price:       input.Price,
		Description: input.Description,
	}
	found, err := s.repo.UpdateByID(input.ID, updates)
	if err != nil {
		return err
	}
	if !found {
		return ErrNotFound
	}
	return nil
}
