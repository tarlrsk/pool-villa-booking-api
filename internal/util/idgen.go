package util

import (
	"crypto/rand"
	"math/big"
	"time"
)

// Unambiguous charset — excludes 0/O and 1/I to avoid confusion.
const charset = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"

// GenerateBookingID mirrors the original generateBookingId() in date-format.ts:
// "BK" + YYMMDD + "-" + 6 random charset characters, e.g. BK260801-AB12CD.
func GenerateBookingID() string {
	datePart := time.Now().Format("060102")
	random := make([]byte, 6)
	for i := range random {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			// crypto/rand failures are effectively unrecoverable on this host;
			// fall back to a fixed character rather than crash a booking request.
			random[i] = charset[0]
			continue
		}
		random[i] = charset[n.Int64()]
	}
	return "BK" + datePart + "-" + string(random)
}
