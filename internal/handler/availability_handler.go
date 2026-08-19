package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /api/availability — two modes, mirrors availability/route.ts:
//   - from/to: unavailable dates for calendar highlighting
//   - checkin/checkout: availability + calculated price for a specific stay
func (h *Handler) GetAvailability(c *gin.Context) {
	checkin := c.Query("checkin")
	checkout := c.Query("checkout")
	from := c.Query("from")
	to := c.Query("to")

	if from != "" && to != "" {
		dates, err := h.AvailabilityGetUnavailableDates.Execute(from, to)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load unavailable dates"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"unavailableDates": dates})
		return
	}

	if checkin == "" || checkout == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "checkin and checkout are required"})
		return
	}
	if checkin >= checkout {
		c.JSON(http.StatusBadRequest, gin.H{"error": "checkout must be after checkin"})
		return
	}

	availability, err := h.AvailabilityCheck.Execute(checkin, checkout)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check availability"})
		return
	}
	price, err := h.PricingCalculate.Execute(checkin, checkout)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to calculate price"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"available": availability.Available,
		"reason":    availability.Reason,
		"price":     price,
	})
}
