package dsupdate /* import "arnested.dk/go/dsupdate" */

import (
	"testing"
)

func TestNewDsUpdate(t *testing.T) {
	dsu, _ := NewDsUpdate()

	if (*dsu != DsUpdate{}) {
		t.Errorf("NewDsUpdate() without options does not return a pointer to a zero value DsUpdate")
	}
}

func TestNewDsUpdateWithUserId(t *testing.T) {
	dsu, _ := NewDsUpdate(UserId("foo"))

	if dsu.userId != "foo" {
		t.Errorf("Could not initialize with setting user ID")
	}
}

func TestNewDsUpdateWithPassword(t *testing.T) {
	dsu, _ := NewDsUpdate(Password("bar"))

	if dsu.password != "bar" {
		t.Errorf("Could not initialize with setting password")
	}
}

func TestNewDsUpdateWithDomain(t *testing.T) {
	dsu, _ := NewDsUpdate(Domain("eksempel.dk"))

	if dsu.domain != "eksempel.dk" {
		t.Errorf("Could not initialize with setting domain")
	}
}

func TestNewDsUpdateWithMultipleOptions(t *testing.T) {
	dsu, _ := NewDsUpdate(UserId("foo"), Password("bar"), Domain("eksempel.dk"))

	if dsu.userId != "foo" {
		t.Errorf("Could not initialize with setting user ID")
	}

	if dsu.password != "bar" {
		t.Errorf("Could not initialize with setting password")
	}

	if dsu.domain != "eksempel.dk" {
		t.Errorf("Could not initialize with setting domain")
	}

	// Options should be able to be specified in random order.
	dsu2, _ := NewDsUpdate(UserId("foo"), Domain("eksempel.dk"), Password("bar"))

	if dsu2.userId != "foo" {
		t.Errorf("Could not initialize with setting user ID")
	}

	if dsu2.password != "bar" {
		t.Errorf("Could not initialize with setting password")
	}

	if dsu2.domain != "eksempel.dk" {
		t.Errorf("Could not initialize with setting domain")
	}
}
