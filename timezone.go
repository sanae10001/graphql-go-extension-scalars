package scalars

import (
	"errors"
	"time"
)

var (
	TimeZoneWrongInputFormat = errors.New("Scalar.TimeZone: wrong input format")
	TimeZoneWrongInputType   = errors.New("Scalar.TimeZone: wrong input type")
)

func NewTimeZoneFromString(s string) (*TimeZone, error) {
	tz := &TimeZone{}
	loc, err := time.LoadLocation(s)
	if err != nil {
		return nil, err
	}
	tz.Location = *loc
	return tz, nil
}

// A tz string
// See: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List
type TimeZone struct {
	time.Location
}

func (TimeZone) ImplementsGraphQLType(name string) bool {
	return name == "TimeZone"
}

func (t *TimeZone) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		loc, err := time.LoadLocation(input)
		if err != nil {
			return TimeZoneWrongInputFormat
		}
		t.Location = *loc
	default:
		return TimeZoneWrongInputType
	}
	return nil
}

func (t TimeZone) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}
