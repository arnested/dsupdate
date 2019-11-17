# Go library for updating DS records with DK Hostmasters proprietary DS Update protocol

[![Maintainability](https://api.codeclimate.com/v1/badges/e7ce795acaeb8b2a5b61/maintainability)](https://codeclimate.com/github/arnested/go-dsupdate/maintainability)
[![Codecov](https://codecov.io/gh/arnested/go-dsupdate/branch/master/graph/badge.svg)](https://codecov.io/gh/arnested/go-dsupdate)
[![Build Status](https://travis-ci.com/arnested/go-dsupdate.svg?branch=master)](https://travis-ci.com/arnested/go-dsupdate)
[![Release](https://img.shields.io/github/release/arnested/go-dsupdate.svg)](https://github.com/arnested/go-dsupdate/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/arnested.dk/go/dsupdate/)](https://goreportcard.com/report/arnested.dk/go/dsupdate)
[![CLA assistant](https://cla-assistant.io/readme/badge/arnested/go-dsupdate)](https://cla-assistant.io/arnested/go-dsupdate)
[![GoDoc](https://godoc.org/arnested.dk/go/dsupdate?status.svg)](https://godoc.org/arnested.dk/go/dsupdate)

```go
import "arnested.dk/go/dsupdate"
```

Package dsupdate is a library for updating DS records with DK
Hostmasters proprietary DS Update protocol.

DS Update is a proprietary protocol and service developed and offered
by DK Hostmaster as an interface for updating DNSSEC related DS
records associated with a .dk domain name.

The service and protocol is documented at
<https://github.com/DK-Hostmaster/dsu-service-specification>.

This package has functionality to update or delete DS records using
the DS Update protocol.




