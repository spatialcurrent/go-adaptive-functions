// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"reflect"
)

func divide(args []interface{}) (interface{}, error) {

	if len(args) != 2 {
		return nil, &ErrorInvalidArguments{Function: "Divide", Arguments: args}
	}

	a := args[0]
	b := args[1]

	switch av := a.(type) {
	case int:
		switch bv := b.(type) {
		case int:
			return av / bv, nil
		case int64:
			return int64(av) / bv, nil
		case float64:
			return float64(av) / bv, nil
		}
	case int64:
		switch bv := b.(type) {
		case int:
			return av / int64(bv), nil
		case int64:
			return av / bv, nil
		case float64:
			return float64(av) / bv, nil
		}
	case float64:
		switch bv := b.(type) {
		case int:
			return av / float64(bv), nil
		case int64:
			return av / float64(bv), nil
		case float64:
			return av / bv, nil
		}
	}

	return nil, &ErrorInvalidArguments{Function: "Divide", Arguments: args}
}

var Divide = Function{
	Name:    "Divide",
	Aliases: []string{"divide"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.Int, reflect.Int}, Output: reflect.Int},
		Definition{Inputs: []interface{}{reflect.Int, reflect.Int64}, Output: reflect.Int},
		Definition{Inputs: []interface{}{reflect.Int, reflect.Float64}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{reflect.Int64, reflect.Int}, Output: reflect.Int},
		Definition{Inputs: []interface{}{reflect.Int64, reflect.Int64}, Output: reflect.Int},
		Definition{Inputs: []interface{}{reflect.Int64, reflect.Float64}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{reflect.Float64, reflect.Int}, Output: reflect.Int},
		Definition{Inputs: []interface{}{reflect.Float64, reflect.Int64}, Output: reflect.Int},
		Definition{Inputs: []interface{}{reflect.Float64, reflect.Float64}, Output: reflect.Float64},
	},
	Function: divide,
}
