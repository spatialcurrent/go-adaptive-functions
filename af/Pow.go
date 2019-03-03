// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"math"
	"reflect"
)

func pow(args []interface{}) (interface{}, error) {

	if len(args) != 2 {
		return nil, &ErrorInvalidArguments{Function: "Pow", Arguments: args}
	}

	a := args[0]
	b := args[1]

	switch av := a.(type) {
	case int:
		switch bv := b.(type) {
		case int:
			return math.Pow(float64(av), float64(bv)), nil
		case int32:
			return math.Pow(float64(av), float64(bv)), nil
		case int64:
			return math.Pow(float64(av), float64(bv)), nil
		case float64:
			return math.Pow(float64(av), bv), nil
		}
	case int32:
		switch bv := b.(type) {
		case int:
			return math.Pow(float64(av), float64(bv)), nil
		case int32:
			return math.Pow(float64(av), float64(bv)), nil
		case int64:
			return math.Pow(float64(av), float64(bv)), nil
		case float64:
			return math.Pow(float64(av), bv), nil
		}
	case int64:
		switch bv := b.(type) {
		case int:
			return math.Pow(float64(av), float64(bv)), nil
		case int32:
			return math.Pow(float64(av), float64(bv)), nil
		case int64:
			return math.Pow(float64(av), float64(bv)), nil
		case float64:
			return math.Pow(float64(av), bv), nil
		}
	case float64:
		switch bv := b.(type) {
		case int:
			return math.Pow(av, float64(bv)), nil
		case int32:
			return math.Pow(av, float64(bv)), nil
		case int64:
			return math.Pow(av, float64(bv)), nil
		case float64:
			return math.Pow(av, bv), nil
		}
	}

	return nil, &ErrorInvalidArguments{Function: "Pow", Arguments: args}
}

var Pow = Function{
	Name:    "Pow",
	Aliases: []string{"pow"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.Int, reflect.Int}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{reflect.Int, reflect.Int32}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{reflect.Int, reflect.Int64}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{reflect.Int, reflect.Float64}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{reflect.Int32, reflect.Int}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{reflect.Int32, reflect.Int32}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{reflect.Int32, reflect.Int64}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{reflect.Int32, reflect.Float64}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{reflect.Int64, reflect.Int}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{reflect.Int64, reflect.Int32}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{reflect.Int64, reflect.Int64}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{reflect.Int64, reflect.Float64}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{reflect.Float64, reflect.Int}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{reflect.Float64, reflect.Int32}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{reflect.Float64, reflect.Int64}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{reflect.Float64, reflect.Float64}, Output: reflect.Float64},
	},
	Function: pow,
}
