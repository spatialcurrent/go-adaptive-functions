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

func length(args ...interface{}) (interface{}, error) {

	if len(args) != 1 {
		return nil, &ErrInvalidArguments{Function: "Length", Arguments: args}
	}

	t := reflect.TypeOf(args[0])

	if !(t.Kind() == reflect.Array || t.Kind() == reflect.Slice || t.Kind() == reflect.Map || t.Kind() == reflect.String) {
		return nil, &ErrInvalidArguments{Function: "Length", Arguments: args}
	}

	return reflect.ValueOf(args[0]).Len(), nil

}

var Length = Function{
	Name:    "Length",
	Aliases: []string{"len", "length"},
	Definitions: []Definition{
		Definition{
			Inputs: []interface{}{reflect.Array},
			Output: reflect.Int,
		},
		Definition{
			Inputs: []interface{}{reflect.Slice},
			Output: reflect.Int,
		},
		Definition{
			Inputs: []interface{}{reflect.Map},
			Output: reflect.Int,
		},
		Definition{
			Inputs: []interface{}{reflect.String},
			Output: reflect.Int,
		},
	},
	Function: length,
}
