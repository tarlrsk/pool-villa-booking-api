package update

import "errors"

var ErrNotFound = errors.New("custom period not found")

type Input struct {
	ID          string
	StartDate   string
	EndDate     string
	Price       float64
	Description string
}

// Service updates an existing custom period.
type Service interface {
	Execute(input Input) error
}
