package sendbookingconfirmation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type lineAPIMessenger struct {
	channelToken string
	client       *http.Client
}

// NewLineAPIMessenger adapts LINE's push-message endpoint to the Messenger port.
func NewLineAPIMessenger(channelToken string) Messenger {
	return &lineAPIMessenger{
		channelToken: channelToken,
		client:       &http.Client{Timeout: 5 * time.Second},
	}
}

func (m *lineAPIMessenger) Push(to, message string) error {
	payload := map[string]interface{}{
		"to": to,
		"messages": []map[string]string{
			{"type": "text", "text": message},
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.line.me/v2/bot/message/push", bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+m.channelToken)

	resp, err := m.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("line push failed with status %d", resp.StatusCode)
	}
	return nil
}
