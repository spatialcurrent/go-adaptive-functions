// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"reflect"
	"time"
)

func min(args []interface{}) (interface{}, error) {

	if len(args) > 2 {
		return nil, &ErrorInvalidArguments{Function: "Min", Arguments: args}
	}

	if len(args) == 2 {
		a := args[0]
		b := args[1]

		at := reflect.TypeOf(a)
		bt := reflect.TypeOf(b)

		if !bt.AssignableTo(at) {
			return nil, &ErrorInvalidArguments{Function: "Min", Arguments: args}
		}

		switch av := a.(type) {
		case int:
			bv := b.(int)
			if av < bv {
				return a, nil
			}
			return b, nil
		case int64:
			bv := b.(int64)
			if av < bv {
				return a, nil
			}
			return b, nil
		case float64:
			bv := b.(float64)
			if av < bv {
				return a, nil
			}
			return b, nil
		case time.Time:
			bv, ok := b.(time.Time)
			if !ok {
				return nil, &ErrorInvalidArguments{Function: "Min", Arguments: args}
			}
			if av.Before(bv) {
				return av, nil
			}
			return bv, nil
		}

		return nil, &ErrorInvalidArguments{Function: "Min", Arguments: args}
	}

	t := reflect.TypeOf(args[0])
	if !(t.Kind() == reflect.Array || t.Kind() == reflect.Slice) {
		if reflect.ValueOf(args[0]).Len() == 0 {
			return nil, &ErrorInvalidArguments{Function: "Min", Arguments: args}
		}
	}

	switch arr := args[0].(type) {
	case []int:
		v := arr[0]
		for i := 1; i < len(arr); i++ {
			if arr[i] < v {
				v = arr[i]
			}
		}
		return v, nil
	case []int64:
		v := arr[0]
		for i := 1; i < len(arr); i++ {
			if arr[i] < v {
				v = arr[i]
			}
		}
		return v, nil
	case []float64:
		v := arr[0]
		for i := 1; i < len(arr); i++ {
			if arr[i] < v {
				v = arr[i]
			}
		}
		return v, nil
	}

	return nil, &ErrorInvalidArguments{Function: "Min", Arguments: args}
}

var Min = Function{
	Name:    "Min",
	Aliases: []string{"min"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.Int, reflect.Int}, Output: reflect.Int},
		Definition{Inputs: []interface{}{reflect.Int64, reflect.Int64}, Output: reflect.Int64},
		Definition{Inputs: []interface{}{reflect.Float64, reflect.Float64}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{timeType, timeType}, Output: timeType},
		Definition{Inputs: []interface{}{intArrayType}, Output: reflect.Int},
		Definition{Inputs: []interface{}{int64ArrayType}, Output: reflect.Int64},
		Definition{Inputs: []interface{}{float64ArrayType}, Output: reflect.Float64},
	},
	Function: min,
}
