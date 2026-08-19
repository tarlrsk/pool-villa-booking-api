package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/deday-pool-villa/backend/internal/middleware"
	bookingcreate "github.com/deday-pool-villa/backend/internal/service/booking/create"
	"github.com/deday-pool-villa/backend/internal/service/lineauth/verifyidtoken"
)

type createBookingRequest struct {
	Phone    string `json:"phone"`
	Checkin  string `json:"checkin"`
	Checkout string `json:"checkout"`
	Guests   int    `json:"guests"`
	Notes    string `json:"notes"`
}

// POST /api/bookings — customer auth required, mirrors bookings/route.ts.
func (h *Handler) CreateBooking(c *gin.Context) {
	profileVal, _ := c.Get(middleware.LineProfileKey)
	profile, _ := profileVal.(*verifyidtoken.Profile)

	var req createBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if req.Phone == "" || req.Checkin == "" || req.Checkout == "" || req.Guests == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}
	if req.Checkin >= req.Checkout {
		c.JSON(http.StatusBadRequest, gin.H{"error": "checkout must be after checkin"})
		return
	}

	output, err := h.BookingCreate.Execute(bookingcreate.Input{
		LineUserID:  profile.UserID,
		DisplayName: profile.DisplayName,
		Phone:       req.Phone,
		Checkin:     req.Checkin,
		Checkout:    req.Checkout,
		Guests:      req.Guests,
		Notes:       req.Notes,
	})
	if err != nil {
		if unavailable, ok := err.(*bookingcreate.UnavailableError); ok {
			reason := unavailable.Reason
			if reason == "" {
				reason = "Dates not available"
			}
			c.JSON(http.StatusConflict, gin.H{"error": reason})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create booking"})
		return
	}

	// Fire-and-forget notification, matching the original .catch(console.error) behavior.
	go func() {
		if err := h.LineNotifySendBookingConfirmation.Execute(output.Booking); err != nil {
			log.Printf("line notify failed for booking %s: %v", output.Booking.ID, err)
		}
	}()

	c.JSON(http.StatusCreated, gin.H{"booking": output.Booking, "price": output.Price})
}

// GET /api/user/bookings — customer auth required, mirrors user/bookings/route.ts.
func (h *Handler) GetUserBookings(c *gin.Context) {
	profileVal, _ := c.Get(middleware.LineProfileKey)
	profile, _ := profileVal.(*verifyidtoken.Profile)

	bookings, err := h.BookingListForUser.Execute(profile.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load bookings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bookings": bookings})
}
