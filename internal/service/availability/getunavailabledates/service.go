package getunavailabledates

import (
	"time"

	portgetunavailabledates "github.com/deday-pool-villa/backend/internal/port/availability/getunavailabledates"
)

const dateLayout = "2006-01-02"

type service struct {
	repo portgetunavailabledates.Repository
}

// New builds the unavailable-dates lookup service.
func New(repo portgetunavailabledates.Repository) Service {
	return &service{repo: repo}
}

// Execute mirrors getUnavailableDates(): every individual date in [from, to)
// that is either explicitly blocked or covered by a pending/confirmed booking.
func (s *service) Execute(from, to string) ([]string, error) {
	set := make(map[string]struct{})

	blocked, err := s.repo.FindBlockedDatesInRange(from, to)
	if err != nil {
		return nil, err
	}
	for _, b := range blocked {
		set[b.Date] = struct{}{}
	}

	bookings, err := s.repo.FindOverlappingBookings(from, to)
	if err != nil {
		return nil, err
	}

	for _, bk := range bookings {
		start, err1 := time.Parse(dateLayout, bk.CheckIn)
		end, err2 := time.Parse(dateLayout, bk.CheckOut)
		if err1 != nil || err2 != nil {
			continue
		}
		for cur := start; cur.Before(end); cur = cur.AddDate(0, 0, 1) {
			dateStr := cur.Format(dateLayout)
			if dateStr >= from && dateStr < to {
				set[dateStr] = struct{}{}
			}
		}
	}

	dates := make([]string, 0, len(set))
	for d := range set {
		dates = append(dates, d)
	}
	return dates, nil
}
