package domain

import "time"

type BlockedDate struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Date      string    `gorm:"type:varchar(10);uniqueIndex;not null" json:"date"`
	Reason    string    `json:"reason"`
	CreatedAt time.Time `json:"createdAt"`
}
