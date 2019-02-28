package dsupdate

import (
	"testing"
)

func TestStatusTest(t *testing.T) {
	expectedText := statusText[SubStatusInvalidUserID]
	statusText := StatusText(SubStatusInvalidUserID)
	if statusText != expectedText {
		t.Errorf("Expected status text '%s' but got '%s'", expectedText, statusText)
	}
}
