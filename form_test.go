package dsupdate

import (
	"testing"
)

func TestForm(t *testing.T) {
	dsu, _ := New(Credentials{Domain: "example.dk", UserID: "XX1234-DK", Password: "correcthorsebatterystaple"})
	dsu.Add(DsRecord{KeyTag: 0, Algorithm: 8, DigestType: 2, Digest: "foo"})
	dsu.Add(DsRecord{KeyTag: 0, Algorithm: 8, DigestType: 2, Digest: "bar"})

	enc := dsu.form().Encode()

	if enc != "algorithm1=8&algorithm2=8&digest1=foo&digest2=bar&digest_type1=2&digest_type2=2&domain=example.dk&keytag1=0&keytag2=0&password=correcthorsebatterystaple&userid=XX1234-DK" {
		t.Errorf("Striong: %s", enc)
	}
}
