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
