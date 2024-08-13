# Go library for updating DS records with Punktum.dk's proprietary DS Update protocol

> [!IMPORTANT]
> Deprecated: Punktum.dk has closed the DS-update Service, see
> <https://punktum.dk/artikler/breaking-changes>

[![CLA assistant](https://cla-assistant.io/readme/badge/arnested/dsupdate)](https://cla-assistant.io/arnested/dsupdate)
[![PkgGoDev](https://pkg.go.dev/badge/arnested.dk/go/dsupdate)](https://pkg.go.dev/arnested.dk/go/dsupdate)

```go
import "arnested.dk/go/dsupdate"
```

Package dsupdate is a library for updating DS records with
Punktum.dk's (DK Hostmasters) proprietary DS Update protocol.

DS Update is a proprietary protocol and service developed and offered
by Punktum as an interface for updating DNSSEC related DS records
associated with a .dk domain name.

The service and protocol is documented at
<https://github.com/Punktum-dk/dsu-service-specification>.

This package has functionality to update or delete DS records using
the DS Update protocol.




