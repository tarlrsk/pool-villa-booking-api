package check

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

func (r *pgRepository) FindBlockedDatesInRange(from, to string) ([]domain.BlockedDate, error) {
	var blocked []domain.BlockedDate
	err := r.db.Where("date >= ? AND date < ?", from, to).Find(&blocked).Error
	return blocked, err
}

func (r *pgRepository) CountOverlappingBookings(checkin, checkout string) (int64, error) {
	var overlapping int64
	err := r.db.Model(&domain.Booking{}).
		Where("status IN ?", []domain.BookingStatus{domain.StatusPending, domain.StatusConfirmed}).
		Where("check_in < ? AND check_out > ?", checkout, checkin).
		Count(&overlapping).Error
	return overlapping, err
}
