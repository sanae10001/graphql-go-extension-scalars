package scalars

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUint_UnmarshalGraphQL(t *testing.T) {
	// int
	input1 := math.MaxUint32
	u := new(Uint)
	err := u.UnmarshalGraphQL(input1)
	assert.NoError(t, err)
	assert.Equal(t, *u, Uint(input1))

	// int32
	var input2 int32 = math.MaxInt32
	u = new(Uint)
	err = u.UnmarshalGraphQL(input2)
	assert.NoError(t, err)
	assert.Equal(t, *u, Uint(input2))

	// float64
	var input3 = math.MaxUint32 / 1.0
	u = new(Uint)
	err = u.UnmarshalGraphQL(input3)
	assert.NoError(t, err)
	assert.Equal(t, *u, Uint(input3))

	// UintWrongInputFormat
	errInput1 := math.MaxUint32 + 1
	u = new(Uint)
	err = u.UnmarshalGraphQL(errInput1)
	assert.EqualError(t, err, UintWrongInputFormat.Error())

	var errInput2 int32 = -100
	u = new(Uint)
	err = u.UnmarshalGraphQL(errInput2)
	assert.EqualError(t, err, UintWrongInputFormat.Error())

	var errInput3 = math.MaxUint32 + 0.5
	u = new(Uint)
	err = u.UnmarshalGraphQL(errInput3)
	assert.EqualError(t, err, UintWrongInputFormat.Error())

	// UintWrongInputType
	errInput4 := "4294967299"
	u = new(Uint)
	err = u.UnmarshalGraphQL(errInput4)
	assert.EqualError(t, err, UintWrongInputType.Error())
}

func TestUint_MarshalJSON(t *testing.T) {
	var u Uint = math.MaxUint32

	test := struct {
		Amount Uint `json:"amount"`
	}{Amount: u}

	tBytes, err := json.Marshal(&test)
	assert.NoError(t, err)
	assert.Equal(t, string(tBytes), `{"amount":4294967295}`)
}
