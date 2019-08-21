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

func second(args ...interface{}) (interface{}, error) {

	if len(args) != 1 {
		return nil, &ErrInvalidArguments{Function: "Second", Arguments: args}
	}

	k := reflect.TypeOf(args[0]).Kind()
	if !(k == reflect.Array || k == reflect.Slice || k == reflect.String) {
		return nil, &ErrInvalidArguments{Function: "Second", Arguments: args}
	}
	v := reflect.ValueOf(args[0])

	if v.Len() < 2 {
		return nil, &ErrInvalidArguments{Function: "Second", Arguments: args}
	}

	return v.Index(1).Interface(), nil
}

var Second = Function{
	Name:    "Second",
	Aliases: []string{"second"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{byteSliceType}, Output: reflect.Uint8},
		Definition{Inputs: []interface{}{uint8SliceType}, Output: reflect.Uint8},
		Definition{Inputs: []interface{}{intSliceType}, Output: reflect.Int},
		Definition{Inputs: []interface{}{int64ArrayType}, Output: reflect.Int64},
		Definition{Inputs: []interface{}{float64ArrayType}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{stringSliceType}, Output: reflect.String},
		Definition{Inputs: []interface{}{stringIntMapSliceType}, Output: stringIntMapType},
		Definition{Inputs: []interface{}{interfaceArrayType}, Output: nil},
		Definition{Inputs: []interface{}{reflect.String}, Output: reflect.Uint8},
	},
	Function: second,
}
