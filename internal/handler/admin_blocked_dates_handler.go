package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	blockeddatescreate "github.com/deday-pool-villa/backend/internal/service/admin/blockeddates/create"
	blockeddatesdelete "github.com/deday-pool-villa/backend/internal/service/admin/blockeddates/delete"
)

// GET /api/admin/blocked-dates — admin auth required.
func (h *Handler) AdminListBlockedDates(c *gin.Context) {
	blocked, err := h.AdminBlockedDatesList.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load blocked dates"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"blockedDates": blocked})
}

type createBlockedDateRequest struct {
	Date   string `json:"date" binding:"required"`
	Reason string `json:"reason"`
}

// POST /api/admin/blocked-dates — admin auth required, mirrors admin/block-dates/route.ts.
func (h *Handler) AdminCreateBlockedDate(c *gin.Context) {
	var req createBlockedDateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "date is required"})
		return
	}

	blocked, err := h.AdminBlockedDatesCreate.Execute(blockeddatescreate.Input{Date: req.Date, Reason: req.Reason})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add blocked date"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"blockedDate": blocked})
}

// DELETE /api/admin/blocked-dates/:id — admin auth required.
func (h *Handler) AdminDeleteBlockedDate(c *gin.Context) {
	id := c.Param("id")
	if err := h.AdminBlockedDatesDelete.Execute(id); err != nil {
		if errors.Is(err, blockeddatesdelete.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blocked date not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete blocked date"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
