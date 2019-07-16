package dsupdate

import (
	"testing"
)

func TestSubStatusStringer48x(t *testing.T) {
	expectedText := "User ID not specified"
	statusText := UserIDNotSpecified.String()
	if statusText != expectedText {
		t.Errorf("Expected status text '%s' but got '%s'", expectedText, statusText)
	}
}

func TestSubStatusStringer49x(t *testing.T) {
	expectedText := "Unknown user ID"
	statusText := UnknownUserID.String()
	if statusText != expectedText {
		t.Errorf("Expected status text '%s' but got '%s'", expectedText, statusText)
	}
}

func TestSubStatusStringer53x(t *testing.T) {
	expectedText := "Authentication failed"
	statusText := AuthenticationFailed.String()
	if statusText != expectedText {
		t.Errorf("Expected status text '%s' but got '%s'", expectedText, statusText)
	}
}

func TestSubStatusStringerUnknown(t *testing.T) {
	expectedText := "SubStatus(100)"
	statusText := SubStatus(100).String()
	if statusText != expectedText {
		t.Errorf("Expected status text '%s' but got '%s'", expectedText, statusText)
	}
}
