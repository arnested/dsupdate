package dsupdate /* import "arnested.dk/go/dsupdate" */

import (
	"errors"

	"github.com/miekg/dns"
)

type DsUpdate struct {
	userID    string
	password  string
	domain    string
	dsRecords []dns.DS
}

func NewDsUpdate(options ...func(*DsUpdate) error) (*DsUpdate, error) {
	dsu := DsUpdate{}

	for _, option := range options {
		err := option(&dsu)

		if err != nil {
			return nil, err
		}
	}

	return &dsu, nil
}

func UserID(userID string) func(*DsUpdate) error {
	return func(dsu *DsUpdate) error {
		return dsu.setUserID(userID)
	}
}

func Password(password string) func(*DsUpdate) error {
	return func(dsu *DsUpdate) error {
		return dsu.setPassword(password)
	}
}

func Domain(domain string) func(*DsUpdate) error {
	return func(dsu *DsUpdate) error {
		return dsu.setDomain(domain)
	}
}

func DS(dsRecords ...dns.DS) func(*DsUpdate) error {
	return func(dsu *DsUpdate) error {
		return dsu.setDsRecords(dsRecords)
	}
}

func (dsu *DsUpdate) setUserID(userID string) error {
	dsu.userID = userID

	return nil
}

func (dsu *DsUpdate) setPassword(password string) error {
	dsu.password = password

	return nil
}

func (dsu *DsUpdate) setDomain(domain string) error {
	dsu.domain = domain

	return nil
}

func (dsu *DsUpdate) setDsRecords(dsRecords []dns.DS) error {
	if len(dsRecords) > 5 {
		return errors.New("Max 5")
	}

	dsu.dsRecords = dsRecords

	return nil
}
