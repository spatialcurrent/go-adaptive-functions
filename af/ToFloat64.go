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

func toFloat64(args []interface{}) (interface{}, error) {
	switch x := args[0].(type) {
	case int:
		return float64(x), nil
	case int8:
		return float64(x), nil
	case int16:
		return float64(x), nil
	case int32:
		return float64(x), nil
	case int64:
		return float64(x), nil
	case float64:
		return x, nil
	}
	return nil, errors.New("invalid arguments for toFloat64")
}

var ToFloat64 = Function{
	Name:    "ToFloat64",
	Aliases: []string{"float64"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{intType}, Output: float64Type},
		Definition{Inputs: []interface{}{int8Type}, Output: float64Type},
		Definition{Inputs: []interface{}{int16Type}, Output: float64Type},
		Definition{Inputs: []interface{}{int32Type}, Output: float64Type},
		Definition{Inputs: []interface{}{int64Type}, Output: float64Type},
		Definition{Inputs: []interface{}{float64Type}, Output: float64Type},
	},
	Function: toFloat64,
}
