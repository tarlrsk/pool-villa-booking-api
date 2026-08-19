package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/deday-pool-villa/backend/internal/domain"
	updatestatus "github.com/deday-pool-villa/backend/internal/service/admin/bookings/updatestatus"
)

// GET /api/admin/bookings?status= — admin auth required.
func (h *Handler) AdminListBookings(c *gin.Context) {
	bookings, err := h.AdminBookingsList.Execute(c.Query("status"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load bookings"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"bookings": bookings})
}

type updateStatusRequest struct {
	Status domain.BookingStatus `json:"status"`
}

// PATCH /api/admin/bookings/:bookingId/status — admin auth required.
func (h *Handler) AdminUpdateBookingStatus(c *gin.Context) {
	bookingID := c.Param("bookingId")

	var req updateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	if err := h.AdminBookingsUpdateStatus.Execute(bookingID, req.Status); err != nil {
		switch {
		case errors.Is(err, updatestatus.ErrInvalidStatus):
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		case errors.Is(err, updatestatus.ErrNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update booking"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
