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

func toInt32(args ...interface{}) (interface{}, error) {
	switch x := args[0].(type) {
	case int:
		return int32(x), nil
	case int8:
		return int32(x), nil
	case int16:
		return int32(x), nil
	case int32:
		return x, nil
	case int64:
		if reflect.ValueOf(int32(0)).OverflowInt(x) {
			return int32(0.0), &ErrorOverflow{Original: x, Target: int32Type}
		}
		return int32(x), nil
	case float32:
		return int32(x), nil
	case float64:
		return int32(x), nil
	case string:
		i, err := strconv.ParseInt(x, 10, 32)
		if err != nil {
			return i, err
		}
		return int32(i), nil
	}
	return nil, &ErrInvalidArguments{Function: "ToInt32", Arguments: args}
}

var ToInt32 = Function{
	Name:    "ToInt32",
	Aliases: []string{"int32"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{intType}, Output: int32Type},
		Definition{Inputs: []interface{}{int8Type}, Output: int32Type},
		Definition{Inputs: []interface{}{int16Type}, Output: int32Type},
		Definition{Inputs: []interface{}{int32Type}, Output: int32Type},
		Definition{Inputs: []interface{}{int64Type}, Output: int32Type},
		Definition{Inputs: []interface{}{float32Type}, Output: int32Type},
		Definition{Inputs: []interface{}{float64Type}, Output: int32Type},
		Definition{Inputs: []interface{}{reflect.String}, Output: int32Type},
	},
	Function: toInt32,
}
