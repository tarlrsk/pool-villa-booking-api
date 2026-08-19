package verifyidtoken

import (
	"strings"

	"github.com/deday-pool-villa/backend/internal/config"
	portverifyidtoken "github.com/deday-pool-villa/backend/internal/port/lineauth/verifyidtoken"
)

type service struct {
	cfg      config.Config
	verifier portverifyidtoken.Verifier
}

// New builds the LINE ID-token verification service.
func New(cfg config.Config, verifier portverifyidtoken.Verifier) Service {
	return &service{cfg: cfg, verifier: verifier}
}

// Execute mirrors verifyLineToken() in line-auth.ts: verify the ID token
// using the channel ID (the numeric prefix of the LIFF ID, before the dash)
// as client_id.
func (s *service) Execute(idToken string) (*Profile, error) {
	if s.cfg.LineLIFFID == "" || idToken == "" {
		return nil, nil
	}
	channelID := strings.Split(s.cfg.LineLIFFID, "-")[0]

	result, err := s.verifier.Verify(idToken, channelID)
	if err != nil || result == nil {
		return nil, err
	}

	return &Profile{UserID: result.Sub, DisplayName: result.Name, PictureURL: result.Picture}, nil
}
