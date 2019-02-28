package dsupdate

import (
	"testing"
)

func TestNew(t *testing.T) {
	dsu, err := New(Credentials{Domain: "example.dk", UserID: "XX1234-DK", Password: "correcthorsebatterystaple"})

	if err != nil {
		t.Error(err)
	}

	if dsu.Domain != "example.dk" {
		t.Errorf("Could not create new with setting domain")
	}

	if dsu.UserID != "XX1234-DK" {
		t.Errorf("Could not create new with setting user ID")
	}

	if dsu.Password != "correcthorsebatterystaple" {
		t.Errorf("Could not create new with setting password")
	}
}

func TestNewWithMissingCredentials(t *testing.T) {
	_, err := New(Credentials{})

	if err == nil {
		t.Error(err)
	}

	_, err = New(Credentials{UserID: "XX1234-DK", Password: "correcthorsebatterystaple"})

	if err == nil {
		t.Error(err)
	}

	_, err = New(Credentials{Domain: "example.dk", Password: "correcthorsebatterystaple"})

	if err == nil {
		t.Error(err)
	}

	_, err = New(Credentials{Domain: "example.dk", UserID: "XX1234-DK"})

	if err == nil {
		t.Error(err)
	}
}

func TestAddDS(t *testing.T) {
	dsu, err := New(Credentials{Domain: "example.dk", UserID: "XX1234-DK", Password: "correcthorsebatterystaple"})

	if err != nil {
		t.Error(err)
	}

	err = dsu.Add(DsRecord{KeyTag: 0, Algorithm: 8, DigestType: 2, Digest: "foo"})

	if err != nil {
		t.Error(err)
	}

	err = dsu.Add(DsRecord{KeyTag: 0, Algorithm: 8, DigestType: 2, Digest: "foo"})

	if err != nil {
		t.Error(err)
	}

	err = dsu.Add(DsRecord{KeyTag: 0, Algorithm: 8, DigestType: 2, Digest: "foo"})

	if err != nil {
		t.Error(err)
	}

	err = dsu.Add(DsRecord{KeyTag: 0, Algorithm: 8, DigestType: 2, Digest: "foo"})

	if err != nil {
		t.Error(err)
	}

	err = dsu.Add(DsRecord{KeyTag: 0, Algorithm: 8, DigestType: 2, Digest: "foo"})

	if err != nil {
		t.Error(err)
	}

	err = dsu.Add(DsRecord{KeyTag: 0, Algorithm: 8, DigestType: 2, Digest: "foo"})

	if err == nil {
		t.Error("Should have errored on adding sixth DS")
	}
}
