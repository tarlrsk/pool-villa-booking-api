package upsert

import (
	"errors"

	"github.com/deday-pool-villa/backend/internal/domain"
)

var ErrEmpty = errors.New("expected an array of {day, price}")

// Service bulk-upserts the fixed 7-row day-rate grid.
type Service interface {
	Execute(rates []domain.DayRate) error
}
