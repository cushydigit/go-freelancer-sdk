package freelancer

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ProjectFrontendStatus string

const (
	ProjectFrontendStatusOpen           ProjectFrontendStatus = "open"
	ProjectFrontendStatusComplete       ProjectFrontendStatus = "complete"
	ProjectFrontendStatusPending        ProjectFrontendStatus = "pending"
	ProjectFrontendStatusDraft          ProjectFrontendStatus = "draft"
	ProjectFrontendStatusWorkInProgress ProjectFrontendStatus = "work_in_progress"
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
	logger       *log.Logger
	httpClient   *http.Client
	rateLimiter  *RateLimiter
	apiToken     string
	baseURL      string
	useRateLimit bool
	Debug        bool

	Services *Services
}

func NewClient(apiToken string) *Client {

	c := &Client{
		logger:       log.Default(),
		httpClient:   &http.Client{},
		baseURL:      BaseAPIMainURL,
		Debug:        false,
		useRateLimit: true,
		rateLimiter:  NewRateLimiter(),
	}

	c.Services = newServices(c)
	return c

}

func (c *Client) debug(format string, v ...any) {
	if c.Debug {
		c.logger.Printf(format, v...)
	}
}

func (c *Client) parseRequest(r *request) (err error) {

	err = r.validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.baseURL, r.endpoint)
	queryString := r.query.Encode()
	header := http.Header{}
	header.Set("freelancer-oauth-v1", c.apiToken)
	header.Set("Content-Type", "application/json")
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	c.debug("full url: %s\n", fullURL)

	r.fullURL = fullURL
	r.header = header
	return nil
}

func (c *Client) callAPI(ctx context.Context, r *request) (data []byte, err error) {
	c.debug("calling api endpoint: %s\n", r.fullURL)
	if err = c.parseRequest(r); err != nil {
		return nil, fmt.Errorf("failed to parse request: %w", err)
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	c.debug("http request: %v\n", req)

	if c.useRateLimit {
		c.rateLimiter.WaitIfNeeded()
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send http request %w", err)
	}

	defer func() {
		err2 := res.Body.Close()
		if err == nil {
			err = err2
		}
	}()

	if c.useRateLimit {
		c.rateLimiter.UpdateFromHeader(res)
	}

	// Read data
	c.debug("http response: %s\n", res.Status)
	data, err = io.ReadAll(res.Body)
	if err != nil {
		c.debug("failed to read http response body: %s\n", err)
		return nil, err
	}

	// Handle error status codes >= 400
	if res.StatusCode >= http.StatusBadRequest {
		c.debug("received bad status code: %s\n", res.Status)

		apiErr := &APIError2{
			Status:   res.Status,
			Response: data,
		}
		// Try to parse structure error
		if json.Valid(data) {
			_ = json.Unmarshal(data, apiErr)
		}
		return nil, apiErr
	}

	return data, nil
}

func execute[T any](ctx context.Context, c *Client, r *request) (T, error) {
	var result T
	data, err := c.callAPI(ctx, r)
	if err != nil {
		return result, fmt.Errorf("api error (%s %s): %w", r.method, r.fullURL, err)
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return result, fmt.Errorf("decode error: %w", err)
	}

	return result, nil
}
