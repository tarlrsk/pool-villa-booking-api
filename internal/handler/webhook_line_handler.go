package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func verifyLineSignature(secret, body, signature string) bool {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(body))
	expected := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expected), []byte(signature))
}

// POST /api/webhook/line — mirrors webhook/line/route.ts: signature-verified,
// currently a no-op stub for future LINE messaging events.
func (h *Handler) LineWebhook(c *gin.Context) {
	rawBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}
	signature := c.GetHeader("x-line-signature")

	if !verifyLineSignature(h.Cfg.LineChannelSecret, string(rawBody), signature) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid signature"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
