package dsupdate

import "testing"

func TestDefaultBaseURL(t *testing.T) {
	t.Parallel()

	client := Client{}

	baseURL := client.BaseURL.String()

	if baseURL != Production.String() {
		t.Errorf("Default base URL is not Production (%s) but '%s'", Production, baseURL)
	}
}

func FuzzForm(f *testing.F) {
	f.Fuzz(func(
		_ *testing.T,
		domain string,
		userID string,
		password string,
		keyTag uint16,
		algorithm uint8,
		digestType uint8,
		digest string,
	) {
		client := Client{
			Domain:   domain,
			UserID:   userID,
			Password: password,
		}

		records := []DsRecord{
			{KeyTag: keyTag, Algorithm: algorithm, DigestType: digestType, Digest: digest},
		}

		client.form(records).Encode()
	})
}
