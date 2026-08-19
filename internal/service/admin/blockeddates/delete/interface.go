package delete

import "errors"

var ErrNotFound = errors.New("blocked date not found")

// Service deletes a blocked date by ID.
type Service interface {
	Execute(id string) error
}
