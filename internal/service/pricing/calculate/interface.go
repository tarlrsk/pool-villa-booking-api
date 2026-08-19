package calculate

type Breakdown struct {
	Date        string  `json:"date"`
	Price       float64 `json:"price"`
	Source      string  `json:"source"` // "custom" | "day-rate"
	Description string  `json:"description"`
}

type Result struct {
	TotalPrice float64     `json:"totalPrice"`
	Nights     int         `json:"nights"`
	Breakdown  []Breakdown `json:"breakdown"`
}

// Service calculates the total price for a stay.
type Service interface {
	Execute(checkin, checkout string) (Result, error)
}
