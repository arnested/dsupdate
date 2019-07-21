package dsupdate

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

// Post the DS records to DK Hostmaster.
func (dsu *DsUpdate) Post(httpClient http.Client) ([]byte, error) {
	resp, err := httpClient.PostForm(dsu.baseURL, dsu.form())

	if err != nil {
		return nil, errors.Wrap(err, "Error creating DS records update request")
	}

	defer func() { _ = resp.Body.Close() }()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusOK {
		return body, nil
	}

	s, ok := subStatus(resp.Header)

	if ok {
		return body, errors.WithMessage(
			errors.WithMessagef(
				s,
				"%d", s,
			),
			"Error updating DS records (DSU substatus)",
		)
	}

	return body, errors.WithMessage(
		errors.WithMessagef(
			errors.New(strconv.Itoa(resp.StatusCode)),
			resp.Status,
		),
		"Error updating DS records (HTTP status)",
	)
}
