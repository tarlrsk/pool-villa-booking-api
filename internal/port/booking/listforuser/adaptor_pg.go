package listforuser

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

func (r *pgRepository) FindByLineUserID(lineUserID string) ([]domain.Booking, error) {
	var bookings []domain.Booking
	err := r.db.Where("line_user_id = ?", lineUserID).Order("created_at desc").Find(&bookings).Error
	return bookings, err
}
