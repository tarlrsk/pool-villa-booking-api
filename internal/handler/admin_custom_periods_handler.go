package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	periodscreate "github.com/deday-pool-villa/backend/internal/service/admin/customperiods/create"
	periodsdelete "github.com/deday-pool-villa/backend/internal/service/admin/customperiods/delete"
	periodsupdate "github.com/deday-pool-villa/backend/internal/service/admin/customperiods/update"
)

// GET /api/admin/custom-periods — admin auth required.
func (h *Handler) AdminListCustomPeriods(c *gin.Context) {
	periods, err := h.AdminCustomPeriodsList.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load custom periods"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"customPeriods": periods})
}

type customPeriodRequest struct {
	StartDate   string  `json:"startDate" binding:"required"`
	EndDate     string  `json:"endDate" binding:"required"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

// POST /api/admin/custom-periods — admin auth required.
func (h *Handler) AdminCreateCustomPeriod(c *gin.Context) {
	var req customPeriodRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "startDate and endDate are required"})
		return
	}

	period, err := h.AdminCustomPeriodsCreate.Execute(periodscreate.Input{
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Price:       req.Price,
		Description: req.Description,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create custom period"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"customPeriod": period})
}

// PUT /api/admin/custom-periods/:id — admin auth required.
func (h *Handler) AdminUpdateCustomPeriod(c *gin.Context) {
	id := c.Param("id")

	var req customPeriodRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "startDate and endDate are required"})
		return
	}

	err := h.AdminCustomPeriodsUpdate.Execute(periodsupdate.Input{
		ID:          id,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Price:       req.Price,
		Description: req.Description,
	})
	if err != nil {
		if errors.Is(err, periodsupdate.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Custom period not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update custom period"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

// DELETE /api/admin/custom-periods/:id — admin auth required.
func (h *Handler) AdminDeleteCustomPeriod(c *gin.Context) {
	id := c.Param("id")
	if err := h.AdminCustomPeriodsDelete.Execute(id); err != nil {
		if errors.Is(err, periodsdelete.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Custom period not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete custom period"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
