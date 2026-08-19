package login

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

func (r *pgRepository) FindByUsername(username string) (domain.Admin, error) {
	var admin domain.Admin
	err := r.db.Where("username = ?", username).First(&admin).Error
	return admin, err
}
