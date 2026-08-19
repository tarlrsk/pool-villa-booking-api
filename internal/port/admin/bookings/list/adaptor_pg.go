package list

import (
	"gorm.io/gorm"

	"github.com/deday-pool-villa/backend/internal/domain"
)

type pgRepository struct{ db *gorm.DB }

// NewPostgresRepository adapts a *gorm.DB to the Repository port.
func NewPostgresRepository(db *gorm.DB) Repository { return &pgRepository{db: db} }

func (r *pgRepository) FindAll(status string) ([]domain.Booking, error) {
	query := r.db.Order("created_at desc")
	if status != "" {
		query = query.Where("status = ?", status)
	}
	var bookings []domain.Booking
	err := query.Find(&bookings).Error
	return bookings, err
}
