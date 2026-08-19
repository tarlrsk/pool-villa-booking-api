package list

import (
	"gorm.io/gorm"

	"github.com/deday-pool-villa/backend/internal/domain"
)

type pgRepository struct{ db *gorm.DB }

// NewPostgresRepository adapts a *gorm.DB to the Repository port.
func NewPostgresRepository(db *gorm.DB) Repository { return &pgRepository{db: db} }

func (r *pgRepository) FindAllOrdered() ([]domain.CustomPeriod, error) {
	var periods []domain.CustomPeriod
	err := r.db.Order("start_date asc").Find(&periods).Error
	return periods, err
}
