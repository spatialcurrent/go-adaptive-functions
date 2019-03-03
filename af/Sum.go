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

func sum(args []interface{}) (interface{}, error) {

	if len(args) != 1 {
		return nil, &ErrorInvalidArguments{Function: "Sum", Arguments: args}
	}

	t := reflect.TypeOf(args[0])
	if !(t.Kind() == reflect.Array || t.Kind() == reflect.Slice) {
		if reflect.ValueOf(args[0]).Len() == 0 {
			return nil, &ErrorInvalidArguments{Function: "Sum", Arguments: args}
		}
	}

	switch arr := args[0].(type) {
	case []uint8:
		v := uint8(0)
		for i := 0; i < len(arr); i++ {
			v += arr[i]
		}
		return v, nil
	case []int:
		v := 0
		for i := 0; i < len(arr); i++ {
			v += arr[i]
		}
		return v, nil
	case []int64:
		v := int64(0)
		for i := 0; i < len(arr); i++ {
			v += arr[i]
		}
		return v, nil
	case []float64:
		v := 0.0
		for i := 0; i < len(arr); i++ {
			v += arr[i]
		}
		return v, nil
	}

	return nil, &ErrorInvalidArguments{Function: "Sum", Arguments: args}
}

var Sum = Function{
	Name:    "Sum",
	Aliases: []string{"sum"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{uint8ArrayType}, Output: reflect.Uint8},
		Definition{Inputs: []interface{}{intArrayType}, Output: reflect.Int},
		Definition{Inputs: []interface{}{int64ArrayType}, Output: reflect.Int64},
		Definition{Inputs: []interface{}{float64ArrayType}, Output: reflect.Float64},
	},
	Function: sum,
}
