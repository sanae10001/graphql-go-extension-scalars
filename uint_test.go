package scalars

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUint_UnmarshalGraphQL(t *testing.T) {
	// int
	input1 := 4294967295
	u := new(Uint)
	err := u.UnmarshalGraphQL(input1)
	assert.NoError(t, err)
	assert.Equal(t, *u, Uint(input1))

	// int32
	var input2 int32 = 2147483647
	u = new(Uint)
	err = u.UnmarshalGraphQL(input2)
	assert.NoError(t, err)
	assert.Equal(t, *u, Uint(input2))

	// error
	errInput1 := 4294967299
	u = new(Uint)
	err = u.UnmarshalGraphQL(errInput1)
	assert.EqualError(t, err, UintWrongInputFormat.Error())

	var errInput2 int32 = -100
	u = new(Uint)
	err = u.UnmarshalGraphQL(errInput2)
	assert.EqualError(t, err, UintWrongInputFormat.Error())

	errInput3 := "4294967299"
	u = new(Uint)
	err = u.UnmarshalGraphQL(errInput3)
	assert.EqualError(t, err, UintWrongInputType.Error())
}

func TestUint_MarshalJSON(t *testing.T) {
	var u Uint = 4294967299
	uBytes, err := json.Marshal(u)
	assert.NoError(t, err)
	assert.Equal(t, string(uBytes), `4294967299`)

	test := struct {
		Amount Uint `json:"amount"`
	}{Amount: u}

	tBytes, err := json.Marshal(&test)
	assert.NoError(t, err)
	assert.Equal(t, string(tBytes), `{"amount":4294967299}`)
}
