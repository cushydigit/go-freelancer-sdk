// Package freelancer provides a Go SDK for interacting with the Freelancer API.
package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	BaseAPIMainURL    = "https://www.freelancer.com/api"
	BaseAPISandBoxURL = "https://api-sandbox.freelancer.com"
)

func (c *Client) GetBaseUrl() string {
	return c.baseURL
}
func (c *Client) SetBaseUrl(url string) {
	c.baseURL = url
}
func (c *Client) SetUseRateLimit(enabled bool) {
	c.useRateLimit = enabled
}

type Client struct {
	httpClient  *http.Client
	logger      *log.Logger
	rateLimiter *RateLimiter

	apiToken     string
	baseURL      string
	useRateLimit bool
	debugMode    bool

	Services *Services
}

type ClientOption func(*Client)

func WithSandBox() ClientOption {
	return func(c *Client) { c.baseURL = BaseAPISandBoxURL }
}

func WithHttpClient(hc *http.Client) ClientOption {
	return func(c *Client) { c.httpClient = hc }
}

func WithDebug(enabled bool) ClientOption {
	return func(c *Client) { c.debugMode = enabled }
}

func NewClient(apiToken string, opts ...ClientOption) *Client {

	c := &Client{
		logger: log.Default(),
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL:      BaseAPIMainURL,
		debugMode:    false,
		useRateLimit: true,
		rateLimiter:  newRateLimiter(),
	}

	for _, opt := range opts {
		opt(c)
	}

	c.Services = newServices(c)
	return c

}

func (c *Client) do(ctx context.Context, method, path string, query url.Values, body io.Reader) ([]byte, error) {

	// Parse Path
	endpoint, err := url.Parse(fmt.Sprintf("%s/%s", c.baseURL, path))
	if err != nil {
		return nil, fmt.Errorf("invalid path: %w", err)
	}

	// Add Query Params
	if query != nil {
		endpoint.RawQuery = query.Encode()
	}

	// Create Request
	req, err := http.NewRequestWithContext(ctx, method, endpoint.String(), body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	// Set headers
	req.Header.Set("freelancer-oauth-v1", c.apiToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "GoFreelancerSDK/1.2 (+github.com/cushydigit/go-freelancer-sdk)")

	// Wait for rate limit
	if c.useRateLimit {
		if err := c.rateLimiter.wait(ctx); err != nil {
			return nil, fmt.Errorf("rate limit: %w", err)
		}
	}

	// Send request
	if c.debugMode {
		c.logger.Printf("--> %s %s", method, endpoint.String())
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("network error: %w", err)
	}
	defer resp.Body.Close()

	// Update rate limit
	if c.useRateLimit {
		c.rateLimiter.updateFromHeader(resp.Header)
	}

	// Handle response
	if c.debugMode {
		c.logger.Printf("<-- %s %s ", resp.Status, endpoint.String())
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Handle errors
	if resp.StatusCode >= 400 {
		apiErr := &APIError{
			StatusCode: resp.StatusCode,
			RawPayload: data,
		}
		// try to parse the JSON error body
		if json.Valid(data) {
			// silent Unmarshaling
			_ = json.Unmarshal(data, apiErr)
		}
		// if the api did not provide a message => fallback
		if apiErr.Message == "" {
			apiErr.Message = http.StatusText(resp.StatusCode)
		}
		return nil, apiErr
	}

	// Handle success
	if c.debugMode {
		c.logger.Printf("<-- %d (%d bytes)", resp.StatusCode, len(data))
	}

	return data, nil
}

func execute[T any](ctx context.Context, c *Client, method, path string, query url.Values, body any) (T, error) {
	var result T

	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return result, err
		}
		bodyReader = bytes.NewReader(b)
	}
	data, err := c.do(ctx, method, path, query, bodyReader)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return result, fmt.Errorf("decode error: %w", err)
	}

	return result, nil
}
