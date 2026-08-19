package check

import (
	portcheck "github.com/deday-pool-villa/backend/internal/port/availability/check"
)

type service struct {
	repo portcheck.Repository
}

// New builds the availability check service.
func New(repo portcheck.Repository) Service {
	return &service{repo: repo}
}

// Execute mirrors checkAvailability() in google-sheets.ts: a stay is
// unavailable if any night in [checkin, checkout) is explicitly blocked, or
// if it overlaps an existing pending/confirmed booking. Overlap rule:
// existing.checkin < newCheckout && existing.checkout > newCheckin.
func (s *service) Execute(checkin, checkout string) (Result, error) {
	blocked, err := s.repo.FindBlockedDatesInRange(checkin, checkout)
	if err != nil {
		return Result{}, err
	}
	if len(blocked) > 0 {
		return Result{Available: false, Reason: "blocked"}, nil
	}

	overlapping, err := s.repo.CountOverlappingBookings(checkin, checkout)
	if err != nil {
		return Result{}, err
	}
	if overlapping > 0 {
		return Result{Available: false, Reason: "booked"}, nil
	}

	return Result{Available: true}, nil
}
