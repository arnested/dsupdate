package dsupdate

import (
	"errors"
)

const (
	// Production base URL for DK Hostmasters DSU service.
	Production = "https://dsu.dk-hostmaster.dk/1.0"
	// Sandbox base URL for DK Hostmasters DSU service.
	Sandbox = "https://dsu-sandbox.dk-hostmaster.dk/1.0"
)

// DsRecord is a DS record
type DsRecord struct {
	KeyTag     uint16
	Algorithm  uint8
	DigestType uint8
	Digest     string
}

// Credentials is domain, user ID and password for authenticating to DK Hostmaster.
type Credentials struct {
	Domain   string
	UserID   string
	Password string
}

// DsUpdate is the main component of the library.
type DsUpdate struct {
	Credentials
	useSandbox bool
	dsRecords  []DsRecord
	baseURL    string
}

// New creates a new DsUpdate.
func New(cred Credentials) (*DsUpdate, error) {
	dsu := &DsUpdate{}
	dsu.Credentials = cred
	dsu.baseURL = Production

	if dsu.Domain == "" {
		return nil, errors.New("No Domain for DK Hostmaster")
	}

	if dsu.UserID == "" {
		return nil, errors.New("No User ID for DK Hostmaster")
	}

	if dsu.Password == "" {
		return nil, errors.New("No Password for DK Hostmaster")
	}

	return dsu, nil
}

// Add a DsRecord to the DsUpdate component.
func (dsu *DsUpdate) Add(ds DsRecord) error {
	if len(dsu.dsRecords) >= 5 {
		return errors.New("Max 5 DS records")
	}

	dsu.dsRecords = append(dsu.dsRecords, ds)

	return nil
}

// BaseURL configures base URL to use for the DSU service
func (dsu *DsUpdate) BaseURL(baseURL string) {
	dsu.baseURL = baseURL
}
