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

func toInt8(args ...interface{}) (interface{}, error) {
	switch x := args[0].(type) {
	case int:
		return int8(x), nil
	case int8:
		return x, nil
	case int16:
		return int8(x), nil
	case int32:
		return int8(x), nil
	case int64:
		return int8(x), nil
	case string:
		i, err := strconv.ParseInt(x, 10, 8)
		if err != nil {
			return i, err
		}
		return int8(i), nil
	}
	return nil, &ErrInvalidArguments{Function: "ToInt8", Arguments: args}
}

var ToInt8 = Function{
	Name:    "ToInt8",
	Aliases: []string{"int8"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{intType}, Output: int8Type},
		Definition{Inputs: []interface{}{int8Type}, Output: int8Type},
		Definition{Inputs: []interface{}{int16Type}, Output: int8Type},
		Definition{Inputs: []interface{}{int32Type}, Output: int8Type},
		Definition{Inputs: []interface{}{int64Type}, Output: int8Type},
		Definition{Inputs: []interface{}{reflect.String}, Output: int8Type},
	},
	Function: toInt8,
}
