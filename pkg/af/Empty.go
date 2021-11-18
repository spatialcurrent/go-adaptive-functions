// =================================================================
//
// Copyright (C) 2021 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"reflect"
)

func empty(args ...interface{}) (interface{}, error) {

	if len(args) != 1 {
		return nil, &ErrInvalidArguments{Function: "Empty", Arguments: args}
	}

	k := reflect.TypeOf(args[0]).Kind()

	if !(k == reflect.Array || k == reflect.Map || k == reflect.Slice || k == reflect.String) {
		return nil, &ErrInvalidArguments{Function: "Empty", Arguments: args}
	}

	return reflect.ValueOf(args[0]).Len() == 0, nil
}

var Empty = Function{
	Name:    "Empty",
	Aliases: []string{"empty"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{byteSliceType}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{uint8SliceType}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{intSliceType}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{int64ArrayType}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{float64ArrayType}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{stringSliceType}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{stringIntMapSliceType}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{interfaceArrayType}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{reflect.String}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{reflect.Map}, Output: reflect.Bool},
	},
	Function: empty,
}
