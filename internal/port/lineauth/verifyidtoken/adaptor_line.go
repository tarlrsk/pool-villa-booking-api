package verifyidtoken

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type lineAPIVerifier struct {
	client *http.Client
}

// NewLineAPIVerifier adapts LINE's OAuth verify endpoint to the Verifier port.
func NewLineAPIVerifier() Verifier {
	return &lineAPIVerifier{client: &http.Client{Timeout: 5 * time.Second}}
}

func (v *lineAPIVerifier) Verify(idToken, channelID string) (*VerifyResult, error) {
	resp, err := v.client.PostForm("https://api.line.me/oauth2/v2.1/verify", url.Values{
		"id_token":  {idToken},
		"client_id": {channelID},
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil
	}

	var data struct {
		Sub     string `json:"sub"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &VerifyResult{Sub: data.Sub, Name: data.Name, Picture: data.Picture}, nil
}
