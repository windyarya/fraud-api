package apis

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/windyarya/backend-final/models"
)

type DiscordWebhookPayload struct {
	Content string `json:"content"`
}

type loggingTransport struct {
    Transport http.RoundTripper
}

func SendNotification(webhookURL string, alert models.Alert) error {
	webFM := fmt.Sprintf("http://ec2-13-215-179-160.ap-southeast-1.compute.amazonaws.com/alerts/%d", alert.ID)
	payload := DiscordWebhookPayload{
		Content: fmt.Sprintf("New Alert Created:\nID: %d\nName: %s\nDetails: %s",
			alert.ID, alert.Name, webFM),
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %v", err)
	}

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
        Transport: &loggingTransport{
            Transport: &http.Transport{
                TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
            },
        },
    }
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK response status: %s", resp.Status)
	}

	return nil
}

func (t *loggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
    log.Printf("Request: %s %s", req.Method, req.URL)
    for name, values := range req.Header {
        for _, value := range values {
            log.Printf("Header: %s: %s", name, value)
        }
    }
    resp, err := t.Transport.RoundTrip(req)
    if err != nil {
        return nil, err
    }
    log.Printf("Response status: %s", resp.Status)
    return resp, nil
}