package dsupdate

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

// Post the DS records to DK Hostmaster.
func (dsu *DsUpdate) Post(httpClient http.Client) ([]byte, Error) {
	resp, err := httpClient.PostForm(dsu.baseURL, dsu.form())

	if err != nil {
		return []byte(err.Error()), dsuError{error: err}
	}

	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		return body, nil
	}

	subStatus := resp.Header.Get("X-DSU")

	if subStatus != "" {
		return nil, subStatusError(resp.StatusCode, subStatus)
	}

	return nil, newErrorf(resp.StatusCode, 0, "DS Upload error: %s", resp.Status)
}

func subStatusError(statusCode int, subStatus string) Error {
	subStatusCode, err := strconv.Atoi(subStatus)

	if err != nil {
		return newErrorf(statusCode, 0, "DS Upload sub-status: %s", subStatus)
	}

	statusText, ok := statusText[subStatusCode]

	if !ok {
		return newErrorf(statusCode, subStatusCode, "DS Upload unknown sub-status: %d", subStatusCode)
	}

	return newErrorf(statusCode, subStatusCode, "DS Upload sub-status: %s", statusText)
}
