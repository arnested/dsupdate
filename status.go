package dsupdate

// DS Upload Sub-status codes.
// See: https://github.com/DK-Hostmaster/dsu-service-specification#http-sub-status-codes)
const (
	SubStatusUserIDNotSpecified                                   = 480
	SubStatusPasswordNotSpecified                                 = 481
	SubStatusMissingAParameter                                    = 482
	SubStatusDomainNameNotSpecified                               = 483
	SubStatusInvalidDomainName                                    = 484
	SubStatusInvalidUserID                                        = 485
	SubStatusInvalidDigestAndDigestTypeCombination                = 486
	SubStatusTheContentsOfAtLeastOneParameterIsSyntacticallyWrong = 487
	SubStatusAtLeastOneDSKeyHasAnInvalidAlgorithm                 = 488
	SubStatusInvalidSequenceOfSets                                = 489
	SubStatusUnknownParameterGiven                                = 495
	SubStatusUnknownUserID                                        = 496
	SubStatusUnknownDomainName                                    = 497
	SubStatusAuthenticationFailed                                 = 531
	SubStatusAuthorizationFailed                                  = 532
	SubStatusAuthenticatingUsingThisPasswordTypeIsNotSupported    = 533
)

var statusText = map[int]string{
	SubStatusUserIDNotSpecified:                                   "Userid not specified",
	SubStatusPasswordNotSpecified:                                 "Password not specified",
	SubStatusMissingAParameter:                                    "Missing a parameter",
	SubStatusDomainNameNotSpecified:                               "Domain name not specified",
	SubStatusInvalidDomainName:                                    "Invalid domain name",
	SubStatusInvalidUserID:                                        "Invalid userid",
	SubStatusInvalidDigestAndDigestTypeCombination:                "Invalid digest and digest_type combination",
	SubStatusTheContentsOfAtLeastOneParameterIsSyntacticallyWrong: "The contents of at least one parameter is syntactically wrong",
	SubStatusAtLeastOneDSKeyHasAnInvalidAlgorithm:                 "At least one DS key has an invalid algorithm",
	SubStatusInvalidSequenceOfSets:                                "Invalid sequence of sets",
	SubStatusUnknownParameterGiven:                                "Unknown parameter given",
	SubStatusUnknownUserID:                                        "Unknown userid",
	SubStatusUnknownDomainName:                                    "Unknown domain name",
	SubStatusAuthenticationFailed:                                 "Authentication failed",
	SubStatusAuthorizationFailed:                                  "Authorization failed",
	SubStatusAuthenticatingUsingThisPasswordTypeIsNotSupported:    "Authenticating using this password type is not supported",
}

// StatusText returns a text for the DSU sub status code. It returns
// the empty string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}
