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

func limit(args []interface{}) (interface{}, error) {

	if len(args) != 2 {
		return nil, &ErrorInvalidArguments{Function: "Limit", Arguments: args}
	}

	if reflect.TypeOf(args[1]).Kind() != reflect.Int {
		return nil, &ErrorInvalidArguments{Function: "Limit", Arguments: args}
	}

	t := reflect.TypeOf(args[0])

	if !(t.Kind() == reflect.Array || t.Kind() == reflect.Slice || t.Kind() == reflect.String) {
		return nil, &ErrorInvalidArguments{Function: "Limit", Arguments: args}
	}

	v := reflect.ValueOf(args[0])
	n := args[1].(int)
	if n > v.Len() {
		return args[0], nil
	}
	return v.Slice(0, n).Interface(), nil
}

var Limit = Function{
	Name:    "Limit",
	Aliases: []string{"limit"},
	Definitions: []Definition{
		Definition{
			Inputs: []interface{}{reflect.Array, reflect.Int},
			Output: reflect.Slice,
		},
		Definition{
			Inputs: []interface{}{reflect.Slice, reflect.Int},
			Output: reflect.Slice,
		},
	},
	Function: limit,
}
