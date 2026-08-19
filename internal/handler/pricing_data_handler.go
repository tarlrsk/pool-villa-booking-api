package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GET /api/pricing-data?from=&to= — day rates + custom periods + unavailable
// dates for the calendar UI. Public, mirrors pricing-data/route.ts.
func (h *Handler) GetPricingData(c *gin.Context) {
	from := c.Query("from")
	if from == "" {
		from = time.Now().Format("2006-01-02")
	}
	to := c.Query("to")
	if to == "" {
		to = time.Now().AddDate(0, 3, 0).Format("2006-01-02")
	}

	dayRates, err := h.AdminDayRatesList.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load day rates"})
		return
	}
	customPeriods, err := h.AdminCustomPeriodsList.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load custom periods"})
		return
	}
	unavailable, err := h.AvailabilityGetUnavailableDates.Execute(from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load unavailable dates"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"dayRates":         dayRates,
		"customPeriods":    customPeriods,
		"unavailableDates": unavailable,
	})
}
