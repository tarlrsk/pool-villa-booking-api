package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/deday-pool-villa/backend/internal/service/adminauth/login"
)

type adminLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// POST /api/admin/login — issues a JWT for the seeded admin account.
func (h *Handler) AdminLogin(c *gin.Context) {
	var req adminLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
		return
	}

	token, err := h.AdminAuthLogin.Execute(req.Username, req.Password)
	if err != nil {
		if errors.Is(err, login.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "login failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
