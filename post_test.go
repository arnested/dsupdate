package dsupdate

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

const (
	connectionClose  = -2
	illegalSubStatus = -1
	noSubStatus      = 0
	unknownSubStatus = 1
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	dsu    *DsUpdate
)

func setup(status int, substatus int) func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if status == connectionClose {
			hj, _ := w.(http.Hijacker)
			conn, _, _ := hj.Hijack()
			conn.Close()

			return
		}

		w.Header().Set("Content-Type", "text/plain")

		switch substatus {
		case noSubStatus:
		case illegalSubStatus:
			w.Header().Set("X-DSU", "foo")
		default:
			w.Header().Set("X-DSU", strconv.Itoa(substatus))
		}

		http.Error(w, "Test server response", status)
	})

	dsu, _ = New(Credentials{Domain: "example.dk", UserID: "XY1234-DK", Password: "correcthorsebatterystaple"})
	dsu.Add(DsRecord{KeyTag: 43930, Algorithm: 8, DigestType: 2, Digest: "E174B66853D0DE1A4E391DFAE924695EB6BF12D28E1A68BDBDB44C4F0D325EA1"})
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
		t.Errorf("Successful post should return OK but failed with error: %s", err)
	}
}

func TestPostAuthFail(t *testing.T) {
	defer setup(http.StatusForbidden, SubStatusAuthenticationFailed)()

	client := http.Client{}

	resp, err := dsu.Post(client)

	if err == nil {
		t.Errorf("Expected error on Authentication failure (with authentication failure sub status) but got response: %s", resp)
	}
}

func TestPostUnknownDSUSubstatus(t *testing.T) {
	defer setup(http.StatusForbidden, unknownSubStatus)()

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

func TestPostStatus(t *testing.T) {
	defer setup(http.StatusForbidden, SubStatusAuthorizationFailed)()

	client := http.Client{}

	_, err := dsu.Post(client)

	if err.Status() != http.StatusForbidden {
		t.Errorf("Expected HTTP status %d but got: %d", http.StatusForbidden, err.Status())
	}

	if err.SubStatus() != SubStatusAuthorizationFailed {
		t.Errorf("Expected DSU sub status %d but got: %d", SubStatusAuthorizationFailed, err.SubStatus())
	}
}
