package dsupdate_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"arnested.dk/go/dsupdate"
)

// This example updates the DS records of the eksempel.dk domain. The
// update is made with at timeout of 5 second specified using a
// http.Client configured with the timeout..
//
//nolint:testableexamples
func Example_update() {
	// Create a client with some fake credentials.
	client := dsupdate.Client{
		Domain:   "eksempel.dk",    // .dk domain name
		UserID:   "ABCD1234-DK",    // Punktum.dk user ID
		Password: "abba4evah",      // Punktum.dk password
		BaseURL:  dsupdate.Sandbox, // If left out defaults to dsupdate.Production
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	// Make a slice of DS records.
	records := []dsupdate.DsRecord{
		{
			KeyTag:     43930,
			Algorithm:  8, // RSA/SHA-256
			DigestType: 2, // SHA-256
			Digest:     "E174B66853D0DE1A4E391DFAE924695EB6BF12D28E1A68BDBDB44C4F0D325EA1",
		},
	}

	ctx := context.Background()

	// Post the new DS record(s) to Punktum.dk.
	resp, err := client.Update(ctx, records)

	// If the update failed and a substatus was returned in the
	// "X-DSU" header the error be of the `SubStatus` type.
	var subStatusErr dsupdate.SubStatus
	if errors.As(err, &subStatusErr) {
		fmt.Printf("Failed with DSU substatus error (%d): %s", subStatusErr, subStatusErr)

		return
	}

	// All other errors will be unspecified error of the error
	// interface.
	if err != nil {
		fmt.Printf("Failed with some error: %s", err)

		return
	}

	// If there was no error returned the update succeeded. `resp`
	// will be the body of whatever was returned from the DS
	// Update service ("Request sent to DSU::Version_1_0 okay").
	fmt.Printf("Succeeded. Punktum.dk responded with the message in the body: %s", resp)
}

// This example deletes existing DS records of the eksempel.dk
// domain. The deletion is made with at timeout of 5 second specified
// using a context with a timeout.
//
//nolint:testableexamples
func Example_delete() {
	// Create a client with some fake credentials.
	client := dsupdate.Client{
		Domain:     "eksempel.dk",    // .dk domain name
		UserID:     "ABCD1234-DK",    // Punktum.dk user ID
		Password:   "abba4evah",      // Punktum.dk password
		BaseURL:    dsupdate.Sandbox, // If left out defaults to dsupdate.Production
		HTTPClient: &http.Client{},   // If left out defaults to http.DefaultClient
	}

	// We'll set a 5 second timeout in the deletion using the
	// context package.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Delete the DS record(s) to Punktum.dk.
	resp, err := client.Delete(ctx)

	// If the update failed and a substatus was returned in the
	// "X-DSU" header the error be of the `SubStatus` type.
	var subStatusErr dsupdate.SubStatus
	if errors.As(err, &subStatusErr) {
		fmt.Printf("Failed with DSU substatus error (%d): %s", subStatusErr, subStatusErr)

		return
	}

	// All other errors will be unspecified error of the error
	// interface.
	if err != nil {
		fmt.Printf("Failed with some error: %s", err)

		return
	}

	// If there was no error returned the delete succeeded. `resp`
	// will be the body of whatever was returned from the DS
	// Update service.
	fmt.Printf("Succeeded. Punktum.dk responded with the message in the body: %s", resp)
}
