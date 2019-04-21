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

func toInt(args []interface{}) (interface{}, error) {
	switch x := args[0].(type) {
	case int:
		return x, nil
	case int8:
		return int(x), nil
	case int16:
		return int(x), nil
	case int32:
		return x, nil
	case int64:
		if reflect.ValueOf(int(0)).OverflowInt(x) {
			return int(0.0), &ErrorOverflow{Original: x, Target: intType}
		}
		return int(x), nil
	case float32:
		return int(x), nil
	case float64:
		return int(x), nil
	case string:
		i, err := strconv.ParseInt(x, 10, 32)
		if err != nil {
			return i, err
		}
		return int32(i), nil
	}
	return nil, &ErrorInvalidArguments{Function: "ToInt", Arguments: args}
}

var ToInt = Function{
	Name:    "ToInt",
	Aliases: []string{"int"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{intType}, Output: intType},
		Definition{Inputs: []interface{}{int8Type}, Output: intType},
		Definition{Inputs: []interface{}{int16Type}, Output: intType},
		Definition{Inputs: []interface{}{int32Type}, Output: intType},
		Definition{Inputs: []interface{}{int64Type}, Output: intType},
		Definition{Inputs: []interface{}{float32Type}, Output: intType},
		Definition{Inputs: []interface{}{float64Type}, Output: intType},
		Definition{Inputs: []interface{}{reflect.String}, Output: intType},
	},
	Function: toInt,
}
