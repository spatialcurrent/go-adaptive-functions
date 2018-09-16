// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"reflect"
)

func getLength(args []interface{}) (interface{}, error) {

	if len(args) != 1 {
		return nil, &ErrorInvalidArguments{Function: "GetLength", Arguments: args}
	}

	t := reflect.TypeOf(args[0])

	if !(t.Kind() == reflect.Array || t.Kind() == reflect.Slice || t.Kind() == reflect.String) {
		return nil, &ErrorInvalidArguments{Function: "GetLength", Arguments: args}
	}

	return reflect.ValueOf(args[0]).Len(), nil

}

var GetLength = Function{
	Name:    "GetLength",
	Aliases: []string{"len"},
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
			Inputs: []interface{}{reflect.String},
			Output: reflect.Int,
		},
	},
	Function: getLength,
}
