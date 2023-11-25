package dsupdate

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Update DS records.
func (c *Client) Update(ctx context.Context, records []DsRecord) ([]byte, error) {
	return c.do(ctx, c.form(records))
}

// Delete DS records.
func (c *Client) Delete(ctx context.Context) ([]byte, error) {
	return c.do(ctx, c.formDelete())
}

func (c *Client) do(ctx context.Context, form url.Values) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.BaseURL.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient().Do(req)
	if (err != nil) || (resp == nil) {
		return nil, fmt.Errorf("failed to contact API: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusOK {
		return body, nil
	}

	s, ok := subStatus(resp.Header)

	if ok {
		return body, s
	}

	//nolint:goerr113
	return body, errors.New(resp.Status)
}
