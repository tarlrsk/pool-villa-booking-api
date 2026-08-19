package getunavailabledates

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

func (r *pgRepository) FindOverlappingBookings(from, to string) ([]domain.Booking, error) {
	var bookings []domain.Booking
	err := r.db.
		Where("status IN ?", []domain.BookingStatus{domain.StatusPending, domain.StatusConfirmed}).
		Where("check_in < ? AND check_out > ?", to, from).
		Find(&bookings).Error
	return bookings, err
}
