package dsupdate_test

import (
	"fmt"
	"net/http"

	"arnested.dk/go/dsupdate"
	"github.com/pkg/errors"
)

func Example() {
	// Create a client with some fake credentials.
	client, _ := dsupdate.New(dsupdate.Credentials{
		Domain:   "eksempel.dk",
		UserID:   "ABCD1234-DK",
		Password: "abba4evah",
	})

	// Let's run this on DK Hostmasters sandbox environment. Can
	// be omitted if you are using `dsudate.Production`.
	client.BaseURL(dsupdate.Sandbox)

	// Add a DS record to the client. Can be called multiple times.
	_ = client.Add(dsupdate.DsRecord{
		KeyTag:     43930,
		Algorithm:  8, // RSA/SHA-256
		DigestType: 2, // SHA-256
		Digest:     "E174B66853D0DE1A4E391DFAE924695EB6BF12D28E1A68BDBDB44C4F0D325EA1",
	})

	// Post the new DS record(s) to DK Hostmaster.
	resp, err := client.Post(http.Client{})

	// If the update failed and a substatus was returned in the
	// "X-DSU" header the error cause will be of the
	// `SubStatusError` type.
	if _, ok := errors.Cause(err).(dsupdate.SubStatusError); ok {
		fmt.Printf("Failed with DSU substatus error: %s", err)

		return
	}

	// All other errors will be unspecified error of the error
	// interface.
	if err != nil {
		fmt.Printf("Failed with some error: %s", err)

		return
	}

	// If there was no error returned the updated
	// succeeded. `resp` will be the body of whatever was returned
	// from the DS Update service ("Request sent to
	// DSU::Version_1_0 okay").
	fmt.Printf("Succeeded. DK Hostmaster responded with the message in the body: %s", resp)
}
