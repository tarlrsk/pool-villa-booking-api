package domain

import "time"

type BookingStatus string

const (
	StatusPending   BookingStatus = "pending"
	StatusConfirmed BookingStatus = "confirmed"
	StatusCancelled BookingStatus = "cancelled"
)

// Booking dates are stored as YYYY-MM-DD strings (not time.Time) so the
// availability/pricing logic can keep using plain lexical string comparison,
// matching the semantics of the original TypeScript implementation.
type Booking struct {
	ID          string        `gorm:"primaryKey;type:varchar(20)" json:"id"`
	LineUserID  string        `gorm:"index;not null" json:"lineUserId"`
	DisplayName string        `json:"displayName"`
	Phone       string        `json:"phone"`
	CheckIn     string        `gorm:"type:varchar(10);index;not null" json:"checkin"`
	CheckOut    string        `gorm:"type:varchar(10);index;not null" json:"checkout"`
	Guests      int           `json:"guests"`
	TotalPrice  float64       `json:"totalPrice"`
	Status      BookingStatus `gorm:"type:varchar(10);index;default:pending" json:"status"`
	Notes       string        `json:"notes"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"-"`
}
