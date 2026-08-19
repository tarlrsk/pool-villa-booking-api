package updatestatus

import (
	"gorm.io/gorm"

	"github.com/deday-pool-villa/backend/internal/domain"
)

type pgRepository struct{ db *gorm.DB }

// NewPostgresRepository adapts a *gorm.DB to the Repository port.
func NewPostgresRepository(db *gorm.DB) Repository { return &pgRepository{db: db} }

func (r *pgRepository) UpdateStatus(bookingID string, status domain.BookingStatus) (bool, error) {
	result := r.db.Model(&domain.Booking{}).Where("id = ?", bookingID).Update("status", status)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}
