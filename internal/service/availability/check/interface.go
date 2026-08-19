package check

type Result struct {
	Available bool   `json:"available"`
	Reason    string `json:"reason,omitempty"`
}

// Service checks whether a stay is available.
type Service interface {
	Execute(checkin, checkout string) (Result, error)
}
