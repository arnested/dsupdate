package dsupdate_test

import (
	"testing"

	"arnested.dk/go/dsupdate"
)

func TestError(t *testing.T) {
	t.Parallel()

	s := dsupdate.SubStatus(480)

	expected := "user ID not specified"

	if actual := s.Error(); expected != actual {
		t.Errorf("Expected %q. Got %q.", expected, actual)
	}
}
