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
	UserIDNotSpecified                                   SubStatus = 480 // user ID not specified
	PasswordNotSpecified                                 SubStatus = 481 // password not specified
	MissingAParameter                                    SubStatus = 482 // missing a parameter
	DomainNameNotSpecified                               SubStatus = 483 // domain name not specified
	InvalidDomainName                                    SubStatus = 484 // invalid domain name
	InvalidUserID                                        SubStatus = 485 // invalid user ID
	InvalidDigestAndDigestTypeCombination                SubStatus = 486 // invalid digest and digest type combination
	TheContentsOfAtLeastOneParameterIsSyntacticallyWrong SubStatus = 487 // the contents of at least one parameter is syntactically wrong
	AtLeastOneDSKeyHasAnInvalidAlgorithm                 SubStatus = 488 // at least one DS key has an invalid algorithm
	InvalidSequenceOfSets                                SubStatus = 489 // invalid sequence of sets
	UnknownParameterGiven                                SubStatus = 495 // unknown parameter given
	UnknownUserID                                        SubStatus = 496 // unknown user ID
	UnknownDomainName                                    SubStatus = 497 // unknown domain name
	AuthenticationFailed                                 SubStatus = 531 // authentication failed
	AuthorizationFailed                                  SubStatus = 532 // authorization failed
	AuthenticatingUsingThisPasswordTypeIsNotSupported    SubStatus = 533 // authenticating using this password type is not supported
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
