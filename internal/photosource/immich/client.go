// Package immich implements photosource.PhotoSource against Immich's
// REST API (ADR-0002). Auth is via shared API key (ADR-0004).
package immich

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gjagils/memorie-server/internal/photosource"
)

const defaultTimeout = 5 * time.Second

type Config struct {
	BaseURL string
	APIKey  string

	HTTPClient *http.Client
}

type Client struct {
	baseURL *url.URL
	apiKey  string
	httpc   *http.Client
}

var _ photosource.PhotoSource = (*Client)(nil)

func New(cfg Config) (*Client, error) {
	if cfg.BaseURL == "" {
		return nil, errors.New("immich: BaseURL is required")
	}
	if cfg.APIKey == "" {
		return nil, errors.New("immich: APIKey is required")
	}
	u, err := url.Parse(strings.TrimRight(cfg.BaseURL, "/"))
	if err != nil {
		return nil, fmt.Errorf("immich: invalid BaseURL %q: %w", cfg.BaseURL, err)
	}
	if u.Scheme == "" || u.Host == "" {
		return nil, fmt.Errorf("immich: BaseURL %q must include scheme and host", cfg.BaseURL)
	}
	hc := cfg.HTTPClient
	if hc == nil {
		hc = &http.Client{Timeout: defaultTimeout}
	}
	return &Client{baseURL: u, apiKey: cfg.APIKey, httpc: hc}, nil
}

// Health verifies that the configured Immich server is reachable and the
// API key is valid. It hits GET /api/users/me — a single call validates
// both connectivity and auth. Returns nil on 200; an error distinguishing
// invalid-key (401/403) from other failures otherwise.
func (c *Client) Health(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL.String()+"/api/users/me", nil)
	if err != nil {
		return fmt.Errorf("immich: build request: %w", err)
	}
	req.Header.Set("x-api-key", c.apiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpc.Do(req)
	if err != nil {
		return fmt.Errorf("immich: request failed: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		_, _ = io.Copy(io.Discard, resp.Body)
		return nil
	case http.StatusUnauthorized, http.StatusForbidden:
		return fmt.Errorf("immich: API key invalid or expired (status %d)", resp.StatusCode)
	default:
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 256))
		return fmt.Errorf("immich: unexpected status %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}
}
