package create

import (
	"github.com/deday-pool-villa/backend/internal/domain"
	portcreate "github.com/deday-pool-villa/backend/internal/port/admin/customperiods/create"
)

type service struct{ repo portcreate.Repository }

// New builds the custom-period creation service.
func New(repo portcreate.Repository) Service { return &service{repo: repo} }

func (s *service) Execute(input Input) (domain.CustomPeriod, error) {
	period := domain.CustomPeriod{
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		Price:       input.Price,
		Description: input.Description,
	}
	if err := s.repo.Create(&period); err != nil {
		return domain.CustomPeriod{}, err
	}
	return period, nil
}
