// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"reflect"
)

func multiply(args []interface{}) (interface{}, error) {

	if len(args) != 2 {
		return nil, &ErrorInvalidArguments{Function: "Multiply", Arguments: args}
	}

	a := args[0]
	b := args[1]

	switch av := a.(type) {
	case int:
		switch bv := b.(type) {
		case int:
			return av * bv, nil
		case int32:
			return av * int(bv), nil
		case int64:
			return int64(av) * bv, nil
		case float64:
			return float64(av) * bv, nil
		}
	case int32:
		switch bv := b.(type) {
		case int:
			return int(av) * bv, nil
		case int32:
			return av * bv, nil
		case int64:
			return int64(av) * bv, nil
		case float64:
			return float64(av) * bv, nil
		}
	case int64:
		switch bv := b.(type) {
		case int:
			return av * int64(bv), nil
		case int32:
			return av * int64(bv), nil
		case int64:
			return av * bv, nil
		case float64:
			return float64(av) * bv, nil
		}
	case float64:
		switch bv := b.(type) {
		case int:
			return av * float64(bv), nil
		case int32:
			return av * float64(bv), nil
		case int64:
			return av * float64(bv), nil
		case float64:
			return av * bv, nil
		}
	}

	return nil, &ErrorInvalidArguments{Function: "Multiply", Arguments: args}
}

var Multiply = Function{
	Name:    "Multiply",
	Aliases: []string{"multiply", "mul"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.Int, reflect.Int}, Output: reflect.Int},
		Definition{Inputs: []interface{}{reflect.Int, reflect.Int32}, Output: reflect.Int},
		Definition{Inputs: []interface{}{reflect.Int, reflect.Int64}, Output: reflect.Int},
		Definition{Inputs: []interface{}{reflect.Int, reflect.Float64}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{reflect.Int32, reflect.Int}, Output: reflect.Int},
		Definition{Inputs: []interface{}{reflect.Int32, reflect.Int32}, Output: reflect.Int32},
		Definition{Inputs: []interface{}{reflect.Int32, reflect.Int64}, Output: reflect.Int64},
		Definition{Inputs: []interface{}{reflect.Int32, reflect.Float64}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{reflect.Int64, reflect.Int}, Output: reflect.Int64},
		Definition{Inputs: []interface{}{reflect.Int64, reflect.Int32}, Output: reflect.Int64},
		Definition{Inputs: []interface{}{reflect.Int64, reflect.Int64}, Output: reflect.Int64},
		Definition{Inputs: []interface{}{reflect.Int64, reflect.Float64}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{reflect.Float64, reflect.Int}, Output: reflect.Int},
		Definition{Inputs: []interface{}{reflect.Float64, reflect.Int32}, Output: reflect.Int},
		Definition{Inputs: []interface{}{reflect.Float64, reflect.Int64}, Output: reflect.Int},
		Definition{Inputs: []interface{}{reflect.Float64, reflect.Float64}, Output: reflect.Float64},
	},
	Function: multiply,
}
