package delete

import "errors"

var ErrNotFound = errors.New("custom period not found")

// Service deletes a custom period by ID.
type Service interface {
	Execute(id string) error
}
