package nylas

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Nylas API documentatino: https://developer.nylas.com/docs/v3/email/#add-labels-to-email-messages

type NylasClientConfig struct {
	ClientID string `long:"client-id" env:"NYLAS_CLIENT_ID" description:"Nylas client ID" required:"true"`
	GrantID  string `long:"grant-id"   env:"NYLAS_GRANT_ID" description:"Nylas Grant ID for Ticketer Admin Account" required:"true"`
	APIKey   string `long:"api-key"   env:"NYLAS_API_KEY" description:"Nylas API key" required:"true"`
	APIURI   string `long:"api-uri"   env:"NYLAS_API_URI" description:"Nylas API URI" required:"true"`
}

type NylasClient struct {
	baseURL string
	grantID string
	apiKey  string
	client  *http.Client
}

func NewNylasClient(cfg NylasClientConfig) *NylasClient {
	return &NylasClient{
		baseURL: cfg.APIURI,
		apiKey:  cfg.APIKey,
		grantID: cfg.GrantID,
		client:  &http.Client{},
	}
}

func (c *NylasClient) ListThreadMessages(ctx context.Context, threadID string) (*ThreadResponse, error) {
	url := fmt.Sprintf("%s/v3/grants/%s/threads/%s", c.baseURL, c.grantID, threadID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var threadResp ThreadResponse
	if err := json.NewDecoder(resp.Body).Decode(&threadResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &threadResp, nil
}

func (c *NylasClient) GetMessages(ctx context.Context, limit int) (*MessagesResponse, error) {
	endpoint := fmt.Sprintf("%s/v3/grants/%s/messages", c.baseURL, c.grantID)

	// Create URL with query parameters
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	q := u.Query()
	q.Set("unread", "true")
	q.Set("limit", fmt.Sprintf("%d", limit))
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var messagesResp MessagesResponse
	if err := json.NewDecoder(resp.Body).Decode(&messagesResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &messagesResp, nil
}

// SendMessage sends a message using the Nylas API
func (c *NylasClient) SendMessage(msg *SendMessageRequest) (*SendMessageResponse, error) {
	endpoint := fmt.Sprintf("%s/v3/grants/%s/messages/send", c.baseURL, c.grantID)

	payload, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var sendResp SendMessageResponse
	if err := json.NewDecoder(resp.Body).Decode(&sendResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &sendResp, nil
}
