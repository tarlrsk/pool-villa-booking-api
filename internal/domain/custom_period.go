package domain

import "time"

type CustomPeriod struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	StartDate   string    `gorm:"type:varchar(10);index;not null" json:"startDate"`
	EndDate     string    `gorm:"type:varchar(10);index;not null" json:"endDate"`
	Price       float64   `gorm:"not null" json:"price"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"-"`
}
