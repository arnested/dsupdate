package dsupdate_test

import (
	"context"
	"net"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"arnested.dk/go/dsupdate"
)

const (
	connectionClose                         = -2
	noSubStatus          dsupdate.SubStatus = 0
	unparseableSubStatus dsupdate.SubStatus = 1
	illegalSubStatus     dsupdate.SubStatus = 2
)

func setup(status int, substatus dsupdate.SubStatus) (dsupdate.Client, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

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
			w.Header().Set("X-DSU", "foo")
		default:
			w.Header().Set("X-DSU", strconv.Itoa(int(substatus)))
		}
		//	time.Sleep(3 * time.Second)
		http.Error(w, "Test server response", status)
	})

	client := dsupdate.Client{
		HTTPClient: &http.Client{
			Timeout: time.Second * 2,
		},
		BaseURL: dsupdate.BaseURL(server.URL),
	}

	return client, func() {
		server.Close()
	}
}

func setupStatusOK() (dsupdate.Client, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte("Request sent to DSU::Version_1_0 okay"))
	})

	client := dsupdate.Client{
		HTTPClient: &http.Client{
			Timeout: time.Second * 2,
		},
		BaseURL: dsupdate.BaseURL(server.URL),
	}

	return client, func() {
		server.Close()
	}
}

func setupSubStatus(substatus dsupdate.SubStatus) (dsupdate.Client, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")

		w.Header().Set("X-DSU", strconv.Itoa(int(substatus)))
		http.Error(w, "Test server response", int(substatus)/100*100)
	})

	client := dsupdate.Client{
		HTTPClient: &http.Client{
			Timeout: time.Second * 2,
		},
		BaseURL: dsupdate.BaseURL(server.URL),
	}

	return client, func() {
		server.Close()
	}
}

func TestUpdateOK(t *testing.T) {
	client, teardown := setupStatusOK()
	defer teardown()

	ctx := context.Background()
	records := []dsupdate.DsRecord{}

	_, err := client.Update(ctx, records)
	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			t.Errorf("Timeoutsss: %s", err)
		} else {
			t.Errorf("Successful post should return OK but failed with error: %s", err)
		}
	}
}

func TestHTTPDefaulClient(t *testing.T) {
	client, teardown := setup(http.StatusOK, noSubStatus)
	defer teardown()

	client.HTTPClient = nil

	ctx := context.Background()
	records := []dsupdate.DsRecord{}

	_, err := client.Update(ctx, records)
	if err != nil {
		t.Errorf("Successful post should return OK but failed with error: %s", err)
	}
}

func TestInvalidURL(t *testing.T) {
	client, teardown := setup(http.StatusOK, noSubStatus)
	defer teardown()

	ctx := context.Background()
	records := []dsupdate.DsRecord{}

	client.BaseURL = "%"

	_, err := client.Update(ctx, records)

	if err == nil {
		t.Errorf("Successful post should return OK but failed with error: %+v", err)
	}
}

func TestUpdateDSUStatuses(t *testing.T) {
	for _, s := range subStatuses {
		s := s
		t.Run(s.String(), func(t *testing.T) {
			client, teardown := setupSubStatus(s)
			defer teardown()

			ctx := context.Background()
			records := []dsupdate.DsRecord{}

			resp, err := client.Update(ctx, records)

			_, ok := err.(dsupdate.SubStatus)

			if err == nil {
				t.Errorf("Expected error but got none. Got response instead: %s", resp)
			}

			if !ok {
				t.Error("Expected error to be of type dsupdate.SubStatus")
			}

			if err != s {
				t.Errorf("Expected error to be '%s', instead got: %s", s, err)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	client, teardown := setup(http.StatusForbidden, unparseableSubStatus)
	defer teardown()

	ctx := context.Background()

	resp, err := client.Delete(ctx)

	if err == nil {
		t.Errorf("Expected error on Authentication failure (with unknown sub status) but got response: %s", resp)
	}
}

func TestUpdateIllegalDSUSubstatus(t *testing.T) {
	client, teardown := setup(http.StatusForbidden, illegalSubStatus)
	defer teardown()

	ctx := context.Background()
	records := []dsupdate.DsRecord{}

	resp, err := client.Update(ctx, records)

	_, ok := err.(dsupdate.SubStatus)

	if !ok {
		t.Error("Expected error to be of type dsupdate.SubStatus")
	}

	if err == nil {
		t.Errorf("Expected error on Authentication failure (with illegal sub status) but got response: %s", resp)
	}
}

func TestUpdateFailWithNoSubStatus(t *testing.T) {
	client, teardown := setup(http.StatusInternalServerError, noSubStatus)
	defer teardown()

	ctx := context.Background()
	records := []dsupdate.DsRecord{}

	resp, err := client.Update(ctx, records)

	if err == nil {
		t.Errorf("Expected error on Internal server error but got response: %s", resp)
	}
}

func TestUpdateConnectionError(t *testing.T) {
	client, teardown := setup(connectionClose, noSubStatus)
	defer teardown()

	ctx := context.Background()
	records := []dsupdate.DsRecord{}

	resp, err := client.Update(ctx, records)

	if err == nil {
		t.Errorf("Expected error on connection close but got response: %s", resp)
	}
}

var subStatuses = []dsupdate.SubStatus{
	dsupdate.UserIDNotSpecified,
	dsupdate.PasswordNotSpecified,
	dsupdate.MissingAParameter,
	dsupdate.DomainNameNotSpecified,
	dsupdate.InvalidDomainName,
	dsupdate.InvalidUserID,
	dsupdate.InvalidDigestAndDigestTypeCombination,
	dsupdate.TheContentsOfAtLeastOneParameterIsSyntacticallyWrong,
	dsupdate.AtLeastOneDSKeyHasAnInvalidAlgorithm,
	dsupdate.InvalidSequenceOfSets,
	dsupdate.UnknownParameterGiven,
	dsupdate.UnknownUserID,
	dsupdate.UnknownDomainName,
	dsupdate.AuthenticationFailed,
	dsupdate.AuthorizationFailed,
	dsupdate.AuthenticatingUsingThisPasswordTypeIsNotSupported,
}
