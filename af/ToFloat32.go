// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"reflect"
	"strconv"
)

func toFloat32(args []interface{}) (interface{}, error) {
	switch x := args[0].(type) {
	case int:
		return float32(x), nil
	case int8:
		return float32(x), nil
	case int16:
		return float32(x), nil
	case int32:
		return float32(x), nil
	case int64:
		return float32(x), nil
	case float32:
		return x, nil
	case float64:
		if reflect.ValueOf(float32(0.0)).OverflowFloat(x) {
			return float32(0.0), &ErrorOverflow{Original: x, Target: float32Type}
		}
		return float32(x), nil
	case string:
		i, err := strconv.ParseFloat(x, 32)
		if err != nil {
			return float32(0.0), err
		}
		return float32(i), nil
	}
	return nil, &ErrorInvalidArguments{Function: "ToFloat32", Arguments: args}
}

var ToFloat32 = Function{
	Name:    "ToFloat32",
	Aliases: []string{"float32"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{intType}, Output: float32Type},
		Definition{Inputs: []interface{}{int8Type}, Output: float32Type},
		Definition{Inputs: []interface{}{int16Type}, Output: float32Type},
		Definition{Inputs: []interface{}{int32Type}, Output: float32Type},
		Definition{Inputs: []interface{}{int64Type}, Output: float32Type},
		Definition{Inputs: []interface{}{float32Type}, Output: float32Type},
		Definition{Inputs: []interface{}{float64Type}, Output: float32Type},
		Definition{Inputs: []interface{}{reflect.String}, Output: float32Type},
	},
	Function: toFloat32,
}
