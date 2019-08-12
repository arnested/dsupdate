package dsupdate

import (
	"net/http"
	"strconv"
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=SubStatus -linecomment

const subStatusHeader = "X-DSU"

// SubStatus is the error specified by the DS Upload service.
// See: https://github.com/DK-Hostmaster/dsu-service-specification#http-sub-status-codes
type SubStatus int16

// DS Upload Sub-status codes.
// See: https://github.com/DK-Hostmaster/dsu-service-specification#http-sub-status-codes
const (
	UserIDNotSpecified                                   SubStatus = 480 // User ID not specified
	PasswordNotSpecified                                 SubStatus = 481 // Password not specified
	MissingAParameter                                    SubStatus = 482 // Missing a parameter
	DomainNameNotSpecified                               SubStatus = 483 // Domain name not specified
	InvalidDomainName                                    SubStatus = 484 // Invalid domain name
	InvalidUserID                                        SubStatus = 485 // Invalid user ID
	InvalidDigestAndDigestTypeCombination                SubStatus = 486 // Invalid digest and digest type combination
	TheContentsOfAtLeastOneParameterIsSyntacticallyWrong SubStatus = 487 // The contents of at least one parameter is syntactically wrong
	AtLeastOneDSKeyHasAnInvalidAlgorithm                 SubStatus = 488 // At least one DS key has an invalid algorithm
	InvalidSequenceOfSets                                SubStatus = 489 // Invalid sequence of sets
	UnknownParameterGiven                                SubStatus = 495 // Unknown parameter given
	UnknownUserID                                        SubStatus = 496 // Unknown user ID
	UnknownDomainName                                    SubStatus = 497 // Unknown domain name
	AuthenticationFailed                                 SubStatus = 531 // Authentication failed
	AuthorizationFailed                                  SubStatus = 532 // Authorization failed
	AuthenticatingUsingThisPasswordTypeIsNotSupported    SubStatus = 533 // Authenticating using this password type is not supported
)

func (e SubStatus) Error() string {
	return e.String()
}

// subStatus retrieves the substatus from a HTTP header set.
func subStatus(h http.Header) (SubStatus, bool) {
	s, err := strconv.Atoi(h.Get(subStatusHeader))

	if err != nil {
		return 0, false
	}

	return SubStatus(s), true
}
