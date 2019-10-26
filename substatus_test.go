package dsupdate_test

import (
	"testing"

	"arnested.dk/go/dsupdate"
)

func TestError(t *testing.T) {
	s := dsupdate.SubStatus(480)

	actual := s.Error()
	expected := "user ID not specified"

	if expected != actual {
		t.Errorf("Expected %q. Got %q.", expected, actual)
	}
}
