package create

import "github.com/deday-pool-villa/backend/internal/domain"

type Input struct {
	StartDate   string
	EndDate     string
	Price       float64
	Description string
}

// Service creates a new custom period.
type Service interface {
	Execute(input Input) (domain.CustomPeriod, error)
}
