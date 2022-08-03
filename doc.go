/*
Package dsupdate is a library for updating DS records with DK
Hostmasters proprietary DS Update protocol.

DS Update is a proprietary protocol and service developed and offered
by DK Hostmaster as an interface for updating DNSSEC related DS
records associated with a .dk domain name.

The service and protocol is documented at
<https://github.com/DK-Hostmaster/dsu-service-specification>.

This package has functionality to update or delete DS records using
the DS Update protocol.
*/
package dsupdate

//go:generate go run github.com/jimmyfrasche/autoreadme -f
