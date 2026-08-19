package list

import (
	"gorm.io/gorm"

	"github.com/deday-pool-villa/backend/internal/domain"
)

type pgRepository struct{ db *gorm.DB }

// NewPostgresRepository adapts a *gorm.DB to the Repository port.
func NewPostgresRepository(db *gorm.DB) Repository { return &pgRepository{db: db} }

func (r *pgRepository) FindAll() ([]domain.DayRate, error) {
	var rates []domain.DayRate
	err := r.db.Find(&rates).Error
	return rates, err
}
