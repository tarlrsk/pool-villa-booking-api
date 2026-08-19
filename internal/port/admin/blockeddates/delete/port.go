package delete

// Repository is the port blocked-date deletion depends on.
type Repository interface {
	DeleteByID(id string) (bool, error)
}
