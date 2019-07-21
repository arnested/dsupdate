package dsupdate

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/pkg/errors"
)

const (
	connectionClose                     = -2
	noSubStatus          SubStatusError = 0
	unparseableSubStatus SubStatusError = 1
	illegalSubStatus     SubStatusError = 2
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	dsu    *DsUpdate
)

func setup(status int, substatus SubStatusError) func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if status == connectionClose {
			hj, _ := w.(http.Hijacker)
			conn, _, _ := hj.Hijack()
			_ = conn.Close()

			return
		}

		w.Header().Set("Content-Type", "text/plain")

		switch substatus {
		case noSubStatus:
		case unparseableSubStatus:
			w.Header().Set(subStatusHeader, "foo")
		default:
			w.Header().Set(subStatusHeader, strconv.Itoa(int(substatus)))

		}

		http.Error(w, "Test server response", status)
	})

	dsu, _ = New(Credentials{Domain: "example.dk", UserID: "XY1234-DK", Password: "correcthorsebatterystaple"})
	_ = dsu.Add(DsRecord{KeyTag: 43930, Algorithm: 8, DigestType: 2, Digest: "E174B66853D0DE1A4E391DFAE924695EB6BF12D28E1A68BDBDB44C4F0D325EA1"})
	dsu.BaseURL(server.URL)

	return func() {
		server.Close()
	}
}

func TestPostOK(t *testing.T) {
	defer setup(http.StatusOK, noSubStatus)()

	client := http.Client{}

	_, err := dsu.Post(client)

	if err != nil {
		t.Errorf("Successful post should return OK but failed with error: %s", errors.Cause(err))
	}
}

func TestPostAuthFail(t *testing.T) {
	defer setup(http.StatusForbidden, AuthenticationFailed)()

	client := http.Client{}

	resp, err := dsu.Post(client)

	if err == nil {
		t.Errorf("Expected error on Authentication failure (with authentication failure sub status) but got response: %s", resp)
	}
}

func TestPostUnknownDSUSubstatus(t *testing.T) {
	defer setup(http.StatusForbidden, unparseableSubStatus)()

	client := http.Client{}

	resp, err := dsu.Post(client)

	if err == nil {
		t.Errorf("Expected error on Authentication failure (with unknown sub status) but got response: %s", resp)
	}
}

func TestPostIllegalDSUSubstatus(t *testing.T) {
	defer setup(http.StatusForbidden, illegalSubStatus)()

	client := http.Client{}

	resp, err := dsu.Post(client)

	if err == nil {
		t.Errorf("Expected error on Authentication failure (with illegal sub status) but got response: %s", resp)
	}
}

func TestPostFailWithNoSubStatus(t *testing.T) {
	defer setup(http.StatusInternalServerError, noSubStatus)()

	client := http.Client{}

	resp, err := dsu.Post(client)

	if err == nil {
		t.Errorf("Expected error on Internal server error but got response: %s", resp)
	}
}

func TestPostConnectionError(t *testing.T) {
	defer setup(connectionClose, noSubStatus)()

	client := http.Client{}

	resp, err := dsu.Post(client)

	if err == nil {
		t.Errorf("Expected error on connection close but got response: %s", resp)
	}
}

var subStatusTests = []struct {
	key       string
	substatus SubStatusError
}{
	{"illegal substatus", illegalSubStatus},
	{UserIDNotSpecified.String(), UserIDNotSpecified},
	{PasswordNotSpecified.String(), PasswordNotSpecified},
	{MissingAParameter.String(), MissingAParameter},
	{DomainNameNotSpecified.String(), DomainNameNotSpecified},
	{InvalidDomainName.String(), InvalidDomainName},
	{InvalidUserID.String(), InvalidUserID},
	{InvalidDigestAndDigestTypeCombination.String(), InvalidDigestAndDigestTypeCombination},
	{TheContentsOfAtLeastOneParameterIsSyntacticallyWrong.String(), TheContentsOfAtLeastOneParameterIsSyntacticallyWrong},
	{AtLeastOneDSKeyHasAnInvalidAlgorithm.String(), AtLeastOneDSKeyHasAnInvalidAlgorithm},
	{InvalidSequenceOfSets.String(), InvalidSequenceOfSets},
	{UnknownParameterGiven.String(), UnknownParameterGiven},
	{UnknownUserID.String(), UnknownUserID},
	{UnknownDomainName.String(), UnknownDomainName},
	{AuthenticationFailed.String(), AuthenticationFailed},
	{AuthorizationFailed.String(), AuthorizationFailed},
	{AuthenticatingUsingThisPasswordTypeIsNotSupported.String(), AuthenticatingUsingThisPasswordTypeIsNotSupported},
}

func TestPostSubStatus(t *testing.T) {
	for _, s := range subStatusTests {
		t.Run(s.key, func(t *testing.T) {
			defer setup(http.StatusInternalServerError, s.substatus)()

			client := http.Client{}

			_, err := dsu.Post(client)
			if errors.Cause(err) != s.substatus {
				t.Errorf("Expected DSU substatus '%s' but got: '%s'", s.substatus, errors.Cause(err))
			}
		})
	}
}

func TestPostSubStatusError(t *testing.T) {
	for _, s := range subStatusTests {
		t.Run(s.key, func(t *testing.T) {
			if s.substatus.Error() != s.substatus.String() {
				t.Errorf("Expected DSU substatus '%s' but got: '%s'", s.substatus.Error(), s.substatus.String())
			}
		})
	}
}
