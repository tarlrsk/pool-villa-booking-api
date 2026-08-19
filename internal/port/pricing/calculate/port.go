package calculate

import "github.com/deday-pool-villa/backend/internal/domain"

// Repository is the port pricing calculation depends on to read rates.
type Repository interface {
	FindAllDayRates() ([]domain.DayRate, error)
	FindAllCustomPeriods() ([]domain.CustomPeriod, error)
}
