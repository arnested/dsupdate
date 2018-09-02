package dsupdate /* import "arnested.dk/go/dsupdate" */

import (
	"testing"

	"github.com/miekg/dns"
)

func TestNewDsUpdate(t *testing.T) {
	_, err := NewDsUpdate()

	if err != nil {
		t.Error(err)
	}
}

func TestNewDsUpdateWithUserID(t *testing.T) {
	dsu, err := NewDsUpdate(UserID("foo"))

	if err != nil {
		t.Error(err)
	}

	if dsu.userID != "foo" {
		t.Errorf("Could not initialize with setting user ID")
	}
}

func TestNewDsUpdateWithPassword(t *testing.T) {
	dsu, err := NewDsUpdate(Password("bar"))

	if err != nil {
		t.Error(err)
	}

	if dsu.password != "bar" {
		t.Errorf("Could not initialize with setting password")
	}
}

func TestNewDsUpdateWithDomain(t *testing.T) {
	dsu, err := NewDsUpdate(Domain("eksempel.dk"))

	if err != nil {
		t.Error(err)
	}

	if dsu.domain != "eksempel.dk" {
		t.Errorf("Could not initialize with setting domain")
	}
}

func TestNewDsUpdateWithMultipleOptions(t *testing.T) {
	dsu, err := NewDsUpdate(UserID("foo"), Password("bar"), Domain("eksempel.dk"))

	if err != nil {
		t.Error(err)
	}

	if dsu.userID != "foo" {
		t.Errorf("Could not initialize with setting user ID")
	}

	if dsu.password != "bar" {
		t.Errorf("Could not initialize with setting password")
	}

	if dsu.domain != "eksempel.dk" {
		t.Errorf("Could not initialize with setting domain")
	}

	// Options should be able to be specified in random order.
	dsu2, err := NewDsUpdate(UserID("foo"), Domain("eksempel.dk"), Password("bar"))

	if err != nil {
		t.Error(err)
	}

	if dsu2.userID != "foo" {
		t.Errorf("Could not initialize with setting user ID")
	}

	if dsu2.password != "bar" {
		t.Errorf("Could not initialize with setting password")
	}

	if dsu2.domain != "eksempel.dk" {
		t.Errorf("Could not initialize with setting domain")
	}
}

func TestNewDsUpdateWithDsRecords(t *testing.T) {
	dsu, err := NewDsUpdate(DS(
		dns.DS{},
		dns.DS{},
		dns.DS{},
		dns.DS{},
	))

	if err != nil {
		t.Error(err)
	}

	if len(dsu.dsRecords) != 4 {
		t.Errorf("Could not initialize with setting 4 DS records")
	}
}

func TestNewDsUpdateWithDsRecordsAsSlice(t *testing.T) {
	dsu, err := NewDsUpdate(DS([]dns.DS{
		dns.DS{},
		dns.DS{},
		dns.DS{},
		dns.DS{},
	}...))

	if err != nil {
		t.Error(err)
	}

	if len(dsu.dsRecords) != 4 {
		t.Errorf("Could not initialize with setting 4 DS records")
	}
}

func TestNewDsUpdateWithTooManyDsRecordsAsSlice(t *testing.T) {
	_, err := NewDsUpdate(DS([]dns.DS{
		dns.DS{},
		dns.DS{},
		dns.DS{},
		dns.DS{},
		dns.DS{},
		dns.DS{},
	}...))

	if err == nil {
		t.Errorf("Did not fail as expected when trying to set 6 DS records")
	}
}

func TestNewDsUpdateWithTooManyDsRecords(t *testing.T) {
	_, err := NewDsUpdate(DS(
		dns.DS{},
		dns.DS{},
		dns.DS{},
		dns.DS{},
		dns.DS{},
		dns.DS{},
	))

	if err == nil {
		t.Errorf("Did not fail as expected when trying to set more than 5 DS records")
	}
}
