package delete

import (
	portdelete "github.com/deday-pool-villa/backend/internal/port/admin/blockeddates/delete"
)

type service struct{ repo portdelete.Repository }

// New builds the blocked-date deletion service.
func New(repo portdelete.Repository) Service { return &service{repo: repo} }

func (s *service) Execute(id string) error {
	found, err := s.repo.DeleteByID(id)
	if err != nil {
		return err
	}
	if !found {
		return ErrNotFound
	}
	return nil
}
