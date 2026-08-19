package create

import "github.com/deday-pool-villa/backend/internal/domain"

type Input struct {
	Date   string
	Reason string
}

// Service creates a new blocked date.
type Service interface {
	Execute(input Input) (domain.BlockedDate, error)
}
