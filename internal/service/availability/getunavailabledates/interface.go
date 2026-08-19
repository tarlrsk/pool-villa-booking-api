package getunavailabledates

// Service returns every individual date in [from, to) that is unavailable.
type Service interface {
	Execute(from, to string) ([]string, error)
}
