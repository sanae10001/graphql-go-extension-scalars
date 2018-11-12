package scalars

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestInt64_UnmarshalGraphQL(t *testing.T) {
	// int
	input := math.MaxInt64
	i := new(Int64)
	err := i.UnmarshalGraphQL(input)
	assert.NoError(t, err)
	assert.Equal(t, *i, Int64(input))

	// int32
	var input2 int32 = math.MaxInt32
	i = new(Int64)
	err = i.UnmarshalGraphQL(input2)
	assert.NoError(t, err)
	assert.Equal(t, *i, Int64(input2))

	// int64
	var input3 int64 = math.MinInt64
	i = new(Int64)
	err = i.UnmarshalGraphQL(input3)
	assert.NoError(t, err)
	assert.Equal(t, *i, Int64(input3))

	// float64
	var input4 = 9223372036854775295.0
	i = new(Int64)
	err = i.UnmarshalGraphQL(input4)
	assert.NoError(t, err)
	assert.Equal(t, *i, Int64(input4))

	// Int64WrongInputFormat
	errInput1 := math.MaxInt64 + 0.5
	i = new(Int64)
	err = i.UnmarshalGraphQL(errInput1)
	assert.EqualError(t, err, Int64WrongInputFormat.Error())

	// Int64WrongInputType
	errInput2 := "1234567"
	i = new(Int64)
	err = i.UnmarshalGraphQL(errInput2)
	assert.EqualError(t, err, Int64WrongInputType.Error())
}

func TestInt64_MarshalJSON(t *testing.T) {
	var i Int64 = math.MaxInt64
	test := struct {
		Amount Int64 `json:"amount"`
	}{Amount:i}

	tBytes, err := json.Marshal(&test)
	assert.NoError(t, err)
	assert.JSONEq(t, string(tBytes), `{"amount":9223372036854775807}`)
}