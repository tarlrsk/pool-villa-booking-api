package delete

// Repository is the port custom-period deletion depends on.
type Repository interface {
	DeleteByID(id string) (bool, error)
}
