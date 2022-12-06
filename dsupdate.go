package dsupdate

import (
	"net/http"
)

// BaseURL is the endpoint of the DS Update service.
type BaseURL string

// DK Hostmasters environments are predefined as constants.
const (
	// Production environment of DK Hostmasters DSU service.
	Production BaseURL = "https://dsu.dk-hostmaster.dk/1.0"
	// Sandbox environment of DK Hostmasters DSU service.
	Sandbox BaseURL = "https://dsu-sandbox.dk-hostmaster.dk/1.0"
)

// String gives you the endpoint as a string. If the BaseURL is not
// set (the zero value) it will give you the Production environment.
func (baseURL BaseURL) String() string {
	if baseURL == "" {
		return string(Production)
	}

	return string(baseURL)
}

// DsRecord is a DS record.
type DsRecord struct {
	KeyTag     uint16
	Algorithm  uint8
	DigestType uint8
	Digest     string
}

// Client for doing updates and deletions.
//
//nolint:lll
type Client struct {
	Domain     string  // .dk domain name, i.e eksempel.dk
	UserID     string  // DK Hostmaster user ID, i.e. ABCD1234-DK
	Password   string  // DK Hostmater password
	BaseURL    BaseURL // DS Update service base URL. You can use constants dsupdate.Production (default) or dsupdate.Sandbox
	HTTPClient *http.Client
}

func (c *Client) httpClient() *http.Client {
	if c.HTTPClient == nil {
		return http.DefaultClient
	}

	return c.HTTPClient
}
