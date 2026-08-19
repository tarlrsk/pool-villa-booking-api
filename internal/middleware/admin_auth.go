package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/deday-pool-villa/backend/internal/service/adminauth/parsetoken"
)

const AdminUsernameKey = "adminUsername"

func AdminAuth(parseToken parsetoken.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		tokenString := strings.TrimPrefix(header, "Bearer ")

		claims, err := parseToken.Execute(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		c.Set(AdminUsernameKey, claims.Username)
		c.Next()
	}
}
