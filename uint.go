package scalars

import (
	"errors"
	"math"
	"strconv"
)

var (
	UintWrongInputFormat = errors.New("Scalar.Uint: wrong input format, expected a unsigned 32-bit integer")
	UintWrongInputType   = errors.New("Scalar.Uint: wrong input type")
)

func NewUintFromInt(u int) *Uint {
	i := Uint(u)
	return &i
}

func NewUint(u uint) *Uint {
	i := Uint(u)
	return &i
}

// An unsigned 32‐bit integer.
// Range: 0 through 4294967295.
type Uint uint

func (Uint) ImplementsGraphQLType(name string) bool {
	return name == "Uint"
}

func (u *Uint) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case int:
		if input < 0 || input > math.MaxUint32 {
			return UintWrongInputFormat
		}
		*u = Uint(input)
	case int32:
		if input < 0 {
			return UintWrongInputFormat
		}
		*u = Uint(input)
	case float64:
		coerced := Uint(input)
		if input < 0 || input > math.MaxUint32 || float64(coerced) != input {
			return UintWrongInputFormat
		}
		*u = coerced
	default:
		return UintWrongInputType
	}
	return nil
}

func (u Uint) MarshalJSON() ([]byte, error) {
	return strconv.AppendUint(nil, uint64(u), 10), nil
}
