package upsert

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/deday-pool-villa/backend/internal/domain"
)

type pgRepository struct{ db *gorm.DB }

// NewPostgresRepository adapts a *gorm.DB to the Repository port.
func NewPostgresRepository(db *gorm.DB) Repository { return &pgRepository{db: db} }

func (r *pgRepository) UpsertAll(rates []domain.DayRate) error {
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "day"}},
		DoUpdates: clause.AssignmentColumns([]string{"price"}),
	}).Create(&rates).Error
}
