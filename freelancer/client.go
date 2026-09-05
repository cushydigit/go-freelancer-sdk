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

	"github.com/cushydigit/go-freelancer-sdk/freelancer/internal/endpoints"
)

func (c *Client) GetBaseUrl() string {
	return c.baseURL
}
func (c *Client) SetBaseUrl(url string) {
	c.baseURL = url
}

type Client struct {
	httpClient *http.Client
	logger     *log.Logger

	apiToken  string
	baseURL   string
	debugMode bool

	Services *Services
}

type ClientOption func(*Client)

func WithSandBox() ClientOption {
	return func(c *Client) { c.baseURL = endpoints.APISandBoxURL }
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
		apiToken:  apiToken,
		baseURL:   endpoints.APIMainURL,
		debugMode: false,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.Services = newServices(c)
	return c

}

func (c *Client) do(ctx context.Context, method, path string, query url.Values, body io.Reader) ([]byte, *ResponseMeta, error) {

	// Parse Path
	endpoint, err := url.Parse(fmt.Sprintf("%s%s", c.baseURL, path))
	if err != nil {
		return nil, nil, fmt.Errorf("invalid path: %w", err)
	}

	// Add Query Params
	if query != nil {
		endpoint.RawQuery = query.Encode()
	}

	// Create Request
	req, err := http.NewRequestWithContext(ctx, method, endpoint.String(), body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create request: %w", err)
	}
	// Set headers
	req.Header.Set("freelancer-oauth-v1", c.apiToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "GoFreelancerSDK/1.4 (+github.com/cushydigit/go-freelancer-sdk)")

	if c.debugMode {
		c.logger.Printf("--> %s %s", method, endpoint.String())
	}

	// Send request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("network error: %w", err)
	}
	defer resp.Body.Close()

	// Handle response
	if c.debugMode {
		c.logger.Printf("<-- %s %s ", resp.Status, endpoint.String())
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read response: %w", err)
	}

	// parse response meta
	meta := parseResponseMeta(resp)

	// Handle errors
	if resp.StatusCode >= 400 {
		apiErr := &APIError{
			StatusCode: resp.StatusCode,
			RawPayload: data,
			Meta:       meta,
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

		return nil, meta, apiErr
	}

	// Handle success
	if c.debugMode {
		c.logger.Printf("<-- %d (%d bytes)", resp.StatusCode, len(data))
	}

	return data, meta, nil
}

func execute[T any](ctx context.Context, c *Client, method, path string, query url.Values, body any) (T, *ResponseMeta, error) {
	var result T

	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return result, nil, err
		}
		bodyReader = bytes.NewReader(b)
	}
	data, meta, err := c.do(ctx, method, path, query, bodyReader)
	if err != nil {
		return result, meta, err
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return result, meta, fmt.Errorf("decode error: %w", err)
	}

	return result, meta, nil
}
