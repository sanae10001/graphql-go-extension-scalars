package scalars

import (
	"errors"
	"net/url"
)

var (
	UriWrongInputFormat = errors.New("Scalar.Uri: wrong input format, expected RFC3986,RFC3987 or RFC6570(level4)")
	UriWrongInputType   = errors.New("Scalar.Uri: wrong input type")
)

func NewUriFromString(s string) (*Uri, error) {
	u := &Uri{}
	oriUrl, err := url.Parse(s)
	if err != nil {
		return nil, err
	}
	u.URL = *oriUrl
	return u, nil
}

// An RFC 3986, RFC 3987, and RFC 6570 (level 4) compliant URI string.
type Uri struct {
	url.URL
}

func (Uri) ImplementsGraphQLType(name string) bool {
	return name == "Uri"
}

func (u *Uri) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		i, err := url.Parse(input)
		if err != nil {
			return UriWrongInputFormat
		}
		u.URL = *i
		return nil
	default:
		return UriWrongInputType
	}
}

func (u Uri) MarshalJSON() ([]byte, error) {
	return []byte(`"` + u.String() + `"`), nil
}
