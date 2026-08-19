package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/deday-pool-villa/backend/internal/config"
	"github.com/deday-pool-villa/backend/internal/service/lineauth/verifyidtoken"
)

const LineProfileKey = "lineProfile"

// LineAuth verifies the Authorization: Bearer <LINE ID token> header against
// LINE's verify endpoint on every request, mirroring the original app's
// stateless approach (no locally-minted session/JWT for customers).
func LineAuth(cfg config.Config, verify verifyidtoken.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractBearerToken(c.GetHeader("Authorization"))
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		if cfg.LineAuthBypass {
			c.Set(LineProfileKey, &verifyidtoken.Profile{UserID: "dev-user", DisplayName: "Dev User"})
			c.Next()
			return
		}

		profile, err := verify.Execute(token)
		if err != nil || profile == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid LINE token"})
			return
		}

		c.Set(LineProfileKey, profile)
		c.Next()
	}
}

func extractBearerToken(authHeader string) string {
	const prefix = "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		return ""
	}
	return authHeader[len(prefix):]
}
