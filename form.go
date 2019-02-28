package dsupdate

import (
	"fmt"
	"net/url"
)

func (dsu *DsUpdate) form() url.Values {
	form := url.Values{}
	form.Set("domain", dsu.Domain)
	form.Set("userid", dsu.UserID)
	form.Set("password", dsu.Password)

	for i, ds := range dsu.dsRecords {
		delta := i + 1
		form.Set(fmt.Sprintf("keytag%d", delta), fmt.Sprintf("%d", ds.KeyTag))
		form.Set(fmt.Sprintf("algorithm%d", delta), fmt.Sprintf("%d", ds.Algorithm))
		form.Set(fmt.Sprintf("digest_type%d", delta), fmt.Sprintf("%d", ds.DigestType))
		form.Set(fmt.Sprintf("digest%d", delta), ds.Digest)
	}

	return form
}
