package scalars

import (
	"errors"
	"time"
)

var (
	DateTimeWrongInputFormat = errors.New("Scalar.DateTime: wrong input format, expected RFC3339")
	DateTimeWrongInputType   = errors.New("Scalar.DateTime: wrong input type")
)

func NewDateTime(t time.Time) *DateTime {
	if t.IsZero() {
		return nil
	}
	d := DateTime{Time: t}
	return &d
}

// An ISO-8601 encoded UTC date string.
// Format: RFC3339, Location: UTC, No nanosecond
// see also: https://en.wikipedia.org/wiki/ISO_8601
type DateTime struct {
	time.Time
}

func (DateTime) ImplementsGraphQLType(name string) bool {
	return name == "DateTime"
}

func (dt *DateTime) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		t, err := time.Parse(time.RFC3339, input)
		if err != nil {
			return DateTimeWrongInputFormat
		}
		dt.Time = convertToNoNanoUTC(t)
		return nil
	default:
		return DateTimeWrongInputType
	}
}

func (dt DateTime) MarshalJSON() ([]byte, error) {
	if dt.Location() != time.UTC {
		dt.Time = dt.UTC()
	}
	if y := dt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("DateTime.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(time.RFC3339)+2)
	b = append(b, '"')
	b = dt.AppendFormat(b, time.RFC3339)
	b = append(b, '"')
	return b, nil
}

func convertToNoNanoUTC(t time.Time) time.Time {
	ut := t.UTC()
	if ut.Nanosecond() == 0 {
		return ut
	} else {
		return time.Date(
			ut.Year(), ut.Month(), ut.Day(),
			ut.Hour(), ut.Minute(), ut.Second(),
			0, ut.Location(),
		)
	}
}
