package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go-backend-labs/01-beginner/github-user-activity/internal/domain"
)

type GitHubClient interface {
	GetUserEvents(username string) ([]domain.Event, error)
}

type githubClient struct {
	httpClient *http.Client
	baseURL    string
}

func NewGitHubClient() GitHubClient {
	return &githubClient{
		httpClient: &http.Client{},
		baseURL:    "https://api.github.com",
	}
}

func (c *githubClient) GetUserEvents(username string) ([]domain.Event, error) {
	url := fmt.Sprintf("%s/users/%s/events", c.baseURL, username)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GitHub API error: %s", string(body))
	}

	var events []domain.Event
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return events, nil
}
