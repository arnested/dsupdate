package dsupdate /* import "arnested.dk/go/dsupdate" */

import "github.com/miekg/dns"

type DsUpdate struct {
	userId   string
	password string
	domain   string
	dsRecord dns.DS
}

func NewDsUpdate(options ...func(*DsUpdate) error) (*DsUpdate, error) {
	dsu := DsUpdate{}

	for _, option := range options {
		option(&dsu)
	}

	return &dsu, nil
}

func UserId(userId string) func(*DsUpdate) error {
	return func(dsu *DsUpdate) error {
		return dsu.setUserId(userId)
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

func (dsu *DsUpdate) setUserId(userId string) error {
	dsu.userId = userId

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
