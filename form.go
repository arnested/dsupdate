package dsupdate

import (
	"net/url"
	"strconv"
)

func (c *Client) form(records []DsRecord) url.Values {
	form := url.Values{}
	form.Set("domain", c.Domain)
	form.Set("userid", c.UserID)
	form.Set("password", c.Password)

	for i, ds := range records {
		delta := i + 1
		keytag := "keytag" + strconv.Itoa(delta)
		algorithm := "algorithm" + strconv.Itoa(delta)
		digestType := "digest_type" + strconv.Itoa(delta)
		digest := "digest" + strconv.Itoa(delta)

		form.Set(keytag, strconv.Itoa(int(ds.KeyTag)))
		form.Set(algorithm, strconv.Itoa(int(ds.Algorithm)))
		form.Set(digestType, strconv.Itoa(int(ds.DigestType)))
		form.Set(digest, ds.Digest)
	}

	return form
}

func (c *Client) formDelete() url.Values {
	const deleteValue = "DELETE_DS"

	form := url.Values{}
	form.Set("domain", c.Domain)
	form.Set("userid", c.UserID)
	form.Set("password", c.Password)

	form.Set("keytag1", deleteValue)
	form.Set("algorithm1", deleteValue)
	form.Set("digest_type1", deleteValue)
	form.Set("digest1", deleteValue)

	return form
}
