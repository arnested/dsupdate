package dsupdate

import "testing"

func TestDefaultBaseURL(t *testing.T) {
	client := Client{}

	baseURL := client.BaseURL.String()

	if baseURL != Production.String() {
		t.Errorf("Default base URL is not Production (%s) but '%s'", Production, baseURL)
	}
}
