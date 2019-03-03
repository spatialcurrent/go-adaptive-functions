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

func toInt16(args []interface{}) (interface{}, error) {
	switch x := args[0].(type) {
	case int:
		return int16(x), nil
	case int8:
		return int16(x), nil
	case int16:
		return x, nil
	case int32:
		return int16(x), nil
	case int64:
		return int16(x), nil
	case string:
		i, err := strconv.ParseInt(x, 10, 16)
		if err != nil {
			return i, err
		}
		return int16(i), nil
	}
	return nil, &ErrorInvalidArguments{Function: "ToInt16", Arguments: args}
}

var ToInt16 = Function{
	Name:    "ToInt16",
	Aliases: []string{"int16"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{intType}, Output: int16Type},
		Definition{Inputs: []interface{}{int8Type}, Output: int16Type},
		Definition{Inputs: []interface{}{int16Type}, Output: int16Type},
		Definition{Inputs: []interface{}{int32Type}, Output: int16Type},
		Definition{Inputs: []interface{}{int64Type}, Output: int16Type},
		Definition{Inputs: []interface{}{reflect.String}, Output: int16Type},
	},
	Function: toInt16,
}
