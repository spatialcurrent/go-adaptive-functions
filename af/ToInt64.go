// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"reflect"
	"strconv"
)

func toInt64(args []interface{}) (interface{}, error) {
	switch x := args[0].(type) {
	case int:
		return int64(x), nil
	case int8:
		return int64(x), nil
	case int16:
		return int64(x), nil
	case int32:
		return int64(x), nil
	case int64:
		return x, nil
	case float32:
		return int32(x), nil
	case float64:
		return int32(x), nil
	case string:
		return strconv.ParseInt(x, 10, 64)
	}
	return nil, &ErrorInvalidArguments{Function: "ToInt64", Arguments: args}
}

var ToInt64 = Function{
	Name:    "ToInt64",
	Aliases: []string{"int64"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{intType}, Output: int64Type},
		Definition{Inputs: []interface{}{int8Type}, Output: int64Type},
		Definition{Inputs: []interface{}{int16Type}, Output: int64Type},
		Definition{Inputs: []interface{}{int32Type}, Output: int64Type},
		Definition{Inputs: []interface{}{int64Type}, Output: int64Type},
		Definition{Inputs: []interface{}{float32Type}, Output: int64Type},
		Definition{Inputs: []interface{}{float64Type}, Output: int64Type},
		Definition{Inputs: []interface{}{reflect.String}, Output: int64Type},
	},
	Function: toInt64,
}
