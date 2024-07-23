package nylas

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// TODO: need to use a request context for this client since it may be used as a part of regular requests

// Nylas API documentatino: https://developer.nylas.com/docs/v3/email/#add-labels-to-email-messages

type NylasClientConfig struct {
	ClientID string `long:"client-id" env:"NYLAS_CLIENT_ID" description:"Nylas client ID" required:"true"`
	GrantID  string `long:"grant-id"   env:"NYLAS_GRANT_ID" default:"me" description:"Nylas Grant ID for Ticketer Admin Account"`
	APIKey   string `long:"api-key"   env:"NYLAS_API_KEY" description:"Nylas API key" required:"true"`
	APIURI   string `long:"api-uri"   env:"NYLAS_API_URI" description:"Nylas API URI" required:"true"`
}

type NylasClient struct {
	httpClient *nylasHTTPClient
}

func NewNylasClient(cfg NylasClientConfig) *NylasClient {
	return &NylasClient{
		httpClient: newNylasHTTPClient(cfg.APIURI, cfg.GrantID, cfg.APIKey),
	}
}

// ListThreadMessages given a threadID, list all messages in the thread
func (c *NylasClient) ListThreadMessages(threadID string) (*MessagesResponse, error) {
	path := fmt.Sprintf("/v3/grants/%s/messages", c.httpClient.grantID)
	query := url.Values{}
	query.Set("thread_id", threadID)
	query.Set("thread_id", threadID)

	responseBody, err := c.httpClient.doRequest("GET", path, query, nil)
	if err != nil {
		return nil, err
	}

	var messagesResp MessagesResponse
	if err := json.Unmarshal(responseBody, &messagesResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &messagesResp, nil
}

// GetFolders get folder information
func (c *NylasClient) GetFolders() (*GetFolersResponse, error) {
	path := fmt.Sprintf("/v3/grants/%s/folders", c.httpClient.grantID)

	responseBody, err := c.httpClient.doRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var resp GetFolersResponse
	if err := json.Unmarshal(responseBody, &resp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &resp, nil
}

// GetThread get thread information
func (c *NylasClient) GetThread(threadID string) (*ThreadResponse, error) {
	path := fmt.Sprintf("/v3/grants/%s/threads/%s", c.httpClient.grantID, threadID)

	responseBody, err := c.httpClient.doRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var messagesResp ThreadResponse
	if err := json.Unmarshal(responseBody, &messagesResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &messagesResp, nil
}

func (c *NylasClient) GetUnreadMessages(limit int) (*MessagesResponse, error) {
	query := url.Values{}
	query.Set("unread", "true")
	query.Set("limit", fmt.Sprintf("%d", limit))

	responseBody, err := c.httpClient.doRequest("GET", fmt.Sprintf("/v3/grants/%s/messages", c.httpClient.grantID), query, nil)
	if err != nil {
		return nil, err
	}

	var messagesResp MessagesResponse
	if err := json.Unmarshal(responseBody, &messagesResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &messagesResp, nil
}

func (c *NylasClient) SendMessage(msg *SendMessageRequest) (*SendMessageResponse, error) {

	// filter out repleis to yourself

	responseBody, err := c.httpClient.doRequest("POST", fmt.Sprintf("/v3/grants/%s/messages/send", c.httpClient.grantID), nil, msg)
	if err != nil {
		return nil, err
	}

	var sendResp SendMessageResponse
	if err := json.Unmarshal(responseBody, &sendResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &sendResp, nil
}

func (c *NylasClient) UpdateMessageReadStatus(messageID string, unread bool) (*UpdateMessageResponse, error) {

	path := fmt.Sprintf("/v3/grants/%s/messages/%s", c.httpClient.grantID, messageID)

	requestBody := struct {
		Unread bool `json:"unread"`
	}{
		Unread: unread,
	}

	responseBody, err := c.httpClient.doRequest("PUT", path, nil, requestBody)
	if err != nil {
		return nil, err
	}

	var messageResp UpdateMessageResponse
	if err := json.Unmarshal(responseBody, &messageResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &messageResp, nil
}
