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

func mean(args []interface{}) (interface{}, error) {

	if len(args) != 1 {
		return nil, &ErrorInvalidArguments{Function: "Mean", Arguments: args}
	}

	t := reflect.TypeOf(args[0])
	if !(t.Kind() == reflect.Array || t.Kind() == reflect.Slice) {
		if reflect.ValueOf(args[0]).Len() == 0 {
			return nil, &ErrorInvalidArguments{Function: "Mean", Arguments: args}
		}
	}

	switch arr := args[0].(type) {
	case []uint8:
		sum := arr[0]
		for i := 1; i < len(arr); i++ {
			sum += arr[i]
		}
		return float64(sum) / float64(len(arr)), nil
	case []int:
		sum := arr[0]
		for i := 1; i < len(arr); i++ {
			sum += arr[i]
		}
		return float64(sum) / float64(len(arr)), nil
	case []int64:
		sum := arr[0]
		for i := 1; i < len(arr); i++ {
			sum += arr[i]
		}
		return float64(sum) / float64(len(arr)), nil
	case []float64:
		sum := arr[0]
		for i := 1; i < len(arr); i++ {
			sum += arr[i]
		}
		return float64(sum) / float64(len(arr)), nil
	}

	return nil, &ErrorInvalidArguments{Function: "Mean", Arguments: args}
}

var Mean = Function{
	Name:    "Mean",
	Aliases: []string{"mean"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{uint8ArrayType}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{intArrayType}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{int64ArrayType}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{float64ArrayType}, Output: reflect.Float64},
	},
	Function: mean,
}
