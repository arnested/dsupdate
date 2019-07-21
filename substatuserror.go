package dsupdate

import (
	"net/http"
	"strconv"
)

//go:generate sh -c "GO111MODULE=off go run golang.org/x/tools/cmd/stringer -type=SubStatusError -linecomment"

const subStatusHeader = "X-DSU"

// SubStatusError is the error specified by the DS Upload service.
// See: https://github.com/DK-Hostmaster/dsu-service-specification#http-sub-status-codes
type SubStatusError int16

// DS Upload Sub-status codes.
// See: https://github.com/DK-Hostmaster/dsu-service-specification#http-sub-status-codes
const (
	UserIDNotSpecified                                   SubStatusError = 480 // User ID not specified
	PasswordNotSpecified                                 SubStatusError = 481 // Password not specified
	MissingAParameter                                    SubStatusError = 482 // Missing a parameter
	DomainNameNotSpecified                               SubStatusError = 483 // Domain name not specified
	InvalidDomainName                                    SubStatusError = 484 // Invalid domain name
	InvalidUserID                                        SubStatusError = 485 // Invalid user ID
	InvalidDigestAndDigestTypeCombination                SubStatusError = 486 // Invalid digest and digest type combination
	TheContentsOfAtLeastOneParameterIsSyntacticallyWrong SubStatusError = 487 // The contents of at least one parameter is syntactically wrong
	AtLeastOneDSKeyHasAnInvalidAlgorithm                 SubStatusError = 488 // At least one DS key has an invalid algorithm
	InvalidSequenceOfSets                                SubStatusError = 489 // Invalid sequence of sets
	UnknownParameterGiven                                SubStatusError = 495 // Unknown parameter given
	UnknownUserID                                        SubStatusError = 496 // Unknown user ID
	UnknownDomainName                                    SubStatusError = 497 // Unknown domain name
	AuthenticationFailed                                 SubStatusError = 531 // Authentication failed
	AuthorizationFailed                                  SubStatusError = 532 // Authorization failed
	AuthenticatingUsingThisPasswordTypeIsNotSupported    SubStatusError = 533 // Authenticating using this password type is not supported
)

func (e SubStatusError) Error() string {
	return e.String()
}

// subStatus from a HTTP header set.
func subStatus(h http.Header) (SubStatusError, bool) {
	s, err := strconv.Atoi(h.Get(subStatusHeader))

	if err != nil {
		return 0, false
	}

	return SubStatusError(s), true
}
