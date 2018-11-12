package scalars

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimeZone_UnmarshalGraphQL(t *testing.T) {
	// string
	tzs := "Africa/Abidjan"
	tz := TimeZone{}
	err := tz.UnmarshalGraphQL(tzs)
	assert.NoError(t, err)
	assert.Equal(t, tzs, tz.String())

	tzs = "Etc/GMT+1"
	tz = TimeZone{}
	err = tz.UnmarshalGraphQL(tzs)
	assert.NoError(t, err)
	assert.Equal(t, tzs, tz.String())

	// TimeZoneWrongInputFormat
	tzs = "Africa/Error"
	tz = TimeZone{}
	err = tz.UnmarshalGraphQL(tzs)
	assert.EqualError(t, err, TimeZoneWrongInputFormat.Error())

	// TimeZoneWrongInputType
	wrongtzs := 123456
	tz = TimeZone{}
	err = tz.UnmarshalGraphQL(wrongtzs)
	assert.EqualError(t, err, TimeZoneWrongInputType.Error())
}

func TestTimeZone_MarshalJSON(t *testing.T) {
	loc := time.UTC
	tz := TimeZone{*loc}
	test := struct {
		TimeZone TimeZone `json:"time_zone"`
	}{tz}
	d, err := json.Marshal(test)
	assert.NoError(t, err)
	assert.JSONEq(t, string(d), `{"time_zone":"UTC"}`)

	// empty
	tz = TimeZone{}
	test = struct {
		TimeZone TimeZone `json:"time_zone"`
	}{TimeZone: tz}
	d, err = json.Marshal(test)
	assert.NoError(t, err)
	assert.JSONEq(t, string(d), `{"time_zone": ""}`)

}
