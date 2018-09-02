# DS Update in Go

[![Maintainability](https://api.codeclimate.com/v1/badges/e7ce795acaeb8b2a5b61/maintainability)](https://codeclimate.com/github/arnested/go-dsupdate/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/e7ce795acaeb8b2a5b61/test_coverage)](https://codeclimate.com/github/arnested/go-dsupdate/test_coverage)
[![Build Status](https://travis-ci.com/arnested/go-dsupdate.svg?branch=master)](https://travis-ci.com/arnested/go-dsupdate)
[![Release](https://img.shields.io/github/release/arnested/go-dsupdate.svg)](https://github.com/arnested/go-dsupdate/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/arnested.dk/go/dsupdate/)](https://goreportcard.com/report/arnested.dk/go/dsupdate)
[![CLA assistant](https://cla-assistant.io/readme/badge/arnested/go-dsupdate)](https://cla-assistant.io/arnested/go-dsupdate)
[![GoDoc](https://godoc.org/arnested.dk/go/dsupdate?status.svg)](https://godoc.org/arnested.dk/go/dsupdate)

# dsupdate
--
    import "arnested.dk/go/dsupdate"

Package dsupdate lorem ipsum ...

## Usage

#### func  DS

```go
func DS(dsRecords ...dns.DS) func(*DsUpdate) error
```

#### func  Domain

```go
func Domain(domain string) func(*DsUpdate) error
```

#### func  Password

```go
func Password(password string) func(*DsUpdate) error
```

#### func  UserID

```go
func UserID(userID string) func(*DsUpdate) error
```

#### type DsUpdate

```go
type DsUpdate struct {
}
```


#### func  NewDsUpdate

```go
func NewDsUpdate(options ...func(*DsUpdate) error) (*DsUpdate, error)
```
