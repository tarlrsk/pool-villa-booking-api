package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/deday-pool-villa/backend/internal/domain"
	dayratesupsert "github.com/deday-pool-villa/backend/internal/service/admin/dayrates/upsert"
)

// GET /api/admin/day-rates — admin auth required.
func (h *Handler) AdminListDayRates(c *gin.Context) {
	rates, err := h.AdminDayRatesList.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load day rates"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"dayRates": rates})
}

// PUT /api/admin/day-rates — bulk upsert the fixed 7-row grid, admin auth required.
func (h *Handler) AdminUpdateDayRates(c *gin.Context) {
	var rates []domain.DayRate
	if err := c.ShouldBindJSON(&rates); err != nil || len(rates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "expected an array of {day, price}"})
		return
	}

	if err := h.AdminDayRatesUpsert.Execute(rates); err != nil {
		if errors.Is(err, dayratesupsert.ErrEmpty) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "expected an array of {day, price}"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update day rates"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
