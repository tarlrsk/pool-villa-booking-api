package list

import (
	"gorm.io/gorm"

	"github.com/deday-pool-villa/backend/internal/domain"
)

type pgRepository struct{ db *gorm.DB }

// NewPostgresRepository adapts a *gorm.DB to the Repository port.
func NewPostgresRepository(db *gorm.DB) Repository { return &pgRepository{db: db} }

func (r *pgRepository) FindAllOrdered() ([]domain.BlockedDate, error) {
	var blocked []domain.BlockedDate
	err := r.db.Order("date asc").Find(&blocked).Error
	return blocked, err
}
