package dsupdate

import (
	"context"
	"errors"
	"io/ioutil"
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
	req, err := http.NewRequest(http.MethodPost, c.BaseURL.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req = req.WithContext(ctx)

	resp, err := c.httpClient().Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusOK {
		return body, nil
	}

	s, ok := subStatus(resp.Header)

	if ok {
		return body, s
	}

	return body, errors.New(resp.Status)
}
