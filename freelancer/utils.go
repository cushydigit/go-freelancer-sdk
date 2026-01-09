package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

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
