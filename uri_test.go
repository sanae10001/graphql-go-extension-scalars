package scalars

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUriFromString(t *testing.T) {
	u, err := NewUriFromString("http://www.ietf.org/rfc/rfc2396.txt")
	assert.NoError(t, err)
	assert.Equal(t, u.Scheme, "http")
	assert.Equal(t, u.Host, "www.ietf.org")
	assert.Equal(t, u.Path, "/rfc/rfc2396.txt")
}

func TestUrl_UnmarshalGraphQL(t *testing.T) {
	// String
	u := Uri{}
	urlUri := "http://www.ietf.org/rfc/rfc2396.txt"
	err := u.UnmarshalGraphQL(urlUri)
	assert.NoError(t, err)
	assert.Equal(t, u.Host, "www.ietf.org")
	assert.Equal(t, u.Scheme, "http")
	assert.Equal(t, u.Path, "/rfc/rfc2396.txt")

	u = Uri{}
	ftpUri := "ftp://ftp.is.co.za/rfc/rfc1808.txt"
	err = u.UnmarshalGraphQL(ftpUri)
	assert.NoError(t, err)
	assert.Equal(t, u.Host, "ftp.is.co.za")
	assert.Equal(t, u.Scheme, "ftp")
	assert.Equal(t, u.Path, "/rfc/rfc1808.txt")

	u = Uri{}
	mailUri := "mailto:John.Doe@example.com"
	err = u.UnmarshalGraphQL(mailUri)
	assert.NoError(t, err)
	assert.Equal(t, u.Scheme, "mailto")
	assert.Equal(t, u.Opaque, "John.Doe@example.com")
}

func TestUri_MarshalJSON(t *testing.T) {
	objUrl := url.URL{
		Host:   "www.ietf.org",
		Scheme: "http",
		Path:   "/rfc/rfc2396.txt",
	}
	uri := Uri{objUrl}
	test := struct {
		Uri Uri `json:"uri"`
	}{uri}
	d, err := json.Marshal(test)
	assert.NoError(t, err)
	assert.Equal(t, string(d), `{"uri":"http://www.ietf.org/rfc/rfc2396.txt"}`)

	// Empty
	uri = Uri{}
	test = struct {
		Uri Uri `json:"uri"`
	}{uri}
	d, err = json.Marshal(test)
	assert.NoError(t, err)
	assert.Equal(t, string(d), `{"uri":""}`)
}
