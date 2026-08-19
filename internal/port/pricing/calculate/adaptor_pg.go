package calculate

import (
	"gorm.io/gorm"

	"github.com/deday-pool-villa/backend/internal/domain"
)

type pgRepository struct {
	db *gorm.DB
}

// NewPostgresRepository adapts a *gorm.DB to the Repository port.
func NewPostgresRepository(db *gorm.DB) Repository {
	return &pgRepository{db: db}
}

func (r *pgRepository) FindAllDayRates() ([]domain.DayRate, error) {
	var rates []domain.DayRate
	err := r.db.Find(&rates).Error
	return rates, err
}

func (r *pgRepository) FindAllCustomPeriods() ([]domain.CustomPeriod, error) {
	var periods []domain.CustomPeriod
	err := r.db.Find(&periods).Error
	return periods, err
}
