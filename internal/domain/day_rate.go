package domain

// DayRate rows use the English weekday name ("Monday".."Sunday") as the
// primary key, mirroring the "Day Rates" sheet from the original app.
type DayRate struct {
	Day   string  `gorm:"primaryKey;type:varchar(10)" json:"day"`
	Price float64 `json:"price"`
}
