package scalars

import (
	"errors"
	"math"
	"strconv"
)

var (
	Int64WrongInputFormat = errors.New("Scalar.Int64: wrong input format, expected a signed 64-bit integers")
	Int64WrongInputType = errors.New("Scalar.Int64: wrong input type")
)

// Int64 is the set of all signed 64-bit integers.
// Range: -9223372036854775808 through 9223372036854775807.
type Int64 int

func (Int64) ImplementsGraphQLType(name string) bool {
	return name == "Int64"
}

func (i *Int64) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case int:
		*i = Int64(input)
	case int32:
		*i = Int64(input)
	case int64:
		*i = Int64(input)
	case float64:
		coerced := Int64(input)
		if input < math.MinInt64 || input > math.MaxInt64 || float64(coerced) != input {
			return Int64WrongInputFormat
		}
		*i = coerced
	default:
		return Int64WrongInputType
	}
	return nil
}

func (i Int64) MarshalJSON() ([]byte, error) {
	return strconv.AppendInt(nil, int64(i), 10), nil
}