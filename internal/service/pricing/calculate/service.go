package calculate

import (
	"time"

	"github.com/deday-pool-villa/backend/internal/domain"
	portcalculate "github.com/deday-pool-villa/backend/internal/port/pricing/calculate"
)

const dateLayout = "2006-01-02"

type service struct {
	repo portcalculate.Repository
}

// New builds the pricing calculation service.
func New(repo portcalculate.Repository) Service {
	return &service{repo: repo}
}

// Execute mirrors calculatePrice() in pricing.ts: iterate each night from
// checkin (inclusive) to checkout (exclusive), custom periods take priority
// over the day-of-week rate, sum to a total.
func (s *service) Execute(checkin, checkout string) (Result, error) {
	dayRates, err := s.repo.FindAllDayRates()
	if err != nil {
		return Result{}, err
	}
	customPeriods, err := s.repo.FindAllCustomPeriods()
	if err != nil {
		return Result{}, err
	}

	dayRateMap := make(map[string]float64, len(dayRates))
	for _, r := range dayRates {
		dayRateMap[r.Day] = r.Price
	}

	start, err := time.Parse(dateLayout, checkin)
	if err != nil {
		return Result{}, err
	}
	end, err := time.Parse(dateLayout, checkout)
	if err != nil {
		return Result{}, err
	}

	breakdown := make([]Breakdown, 0)
	for cur := start; cur.Before(end); cur = cur.AddDate(0, 0, 1) {
		dateStr := cur.Format(dateLayout)

		var custom *domain.CustomPeriod
		for i := range customPeriods {
			p := customPeriods[i]
			if dateStr >= p.StartDate && dateStr < p.EndDate {
				custom = &p
				break
			}
		}

		if custom != nil {
			breakdown = append(breakdown, Breakdown{
				Date:        dateStr,
				Price:       custom.Price,
				Source:      "custom",
				Description: custom.Description,
			})
		} else {
			dayName := cur.Weekday().String()
			breakdown = append(breakdown, Breakdown{
				Date:        dateStr,
				Price:       dayRateMap[dayName],
				Source:      "day-rate",
				Description: dayName,
			})
		}
	}

	total := 0.0
	for _, b := range breakdown {
		total += b.Price
	}

	return Result{TotalPrice: total, Nights: len(breakdown), Breakdown: breakdown}, nil
}
