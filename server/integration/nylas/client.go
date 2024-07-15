package nylas

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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
