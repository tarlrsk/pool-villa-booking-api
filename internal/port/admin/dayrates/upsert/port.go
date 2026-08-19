package upsert

import "github.com/deday-pool-villa/backend/internal/domain"

// Repository is the port day-rates bulk upsert depends on.
type Repository interface {
	UpsertAll(rates []domain.DayRate) error
}
