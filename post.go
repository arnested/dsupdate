package dsupdate

import (
	"io/ioutil"
	"net/http"
)

// Post the DS records to DK Hostmaster.
func (dsu *DsUpdate) Post(httpClient http.Client) ([]byte, Error) {
	resp, err := httpClient.PostForm(dsu.baseURL, dsu.form())

	if err != nil {
		return []byte(err.Error()), dsuError{error: err}
	}

	if resp.StatusCode == http.StatusOK {
		defer func() { _ = resp.Body.Close() }()
		body, _ := ioutil.ReadAll(resp.Body)

		return body, nil
	}

	s, ok := subStatus(resp.Header)

	if ok {
		return nil, newErrorf(resp.StatusCode, int(s), "DS Upload sub-status: %s (%d)", s, s)
	}

	return nil, newErrorf(resp.StatusCode, int(s), "DS Upload error: %s", resp.Status)
}
