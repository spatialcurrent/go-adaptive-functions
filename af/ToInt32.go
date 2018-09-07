// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"github.com/pkg/errors"
)

func toInt32(args []interface{}) (interface{}, error) {
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
		return int32(x), nil
	}
	return nil, errors.New("invalid arguments for toInt32")
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
	},
	Function: toInt32,
}
