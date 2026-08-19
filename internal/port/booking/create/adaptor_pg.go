package create

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

func (r *pgRepository) Create(booking *domain.Booking) error {
	return r.db.Create(booking).Error
}
