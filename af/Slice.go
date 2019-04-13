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

var Slice = Function{
	Name:    "Slice",
	Aliases: []string{"slice"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.Array, intType}, Output: reflect.Slice},
		Definition{Inputs: []interface{}{reflect.Array, intType, intType}, Output: reflect.Slice},
		Definition{Inputs: []interface{}{reflect.Slice, intType}, Output: reflect.Slice},
		Definition{Inputs: []interface{}{reflect.Slice, intType, intType}, Output: reflect.Slice},
		Definition{Inputs: []interface{}{reflect.String, intType}, Output: reflect.String},
		Definition{Inputs: []interface{}{reflect.String, intType, intType}, Output: reflect.String},
	},
	Function: func(args []interface{}) (interface{}, error) {
		if len(args) == 2 {
			if start, ok := args[1].(int); ok {
				v := reflect.ValueOf(args[0])
				switch v.Type().Kind() {
				case reflect.Array, reflect.Slice, reflect.String:
					return v.Slice(start, v.Len()).Interface(), nil
				}
			}
		} else if len(args) == 3 {
			if start, ok := args[1].(int); ok {
				if end, ok := args[2].(int); ok {
					v := reflect.ValueOf(args[0])
					switch v.Type().Kind() {
					case reflect.Array, reflect.Slice, reflect.String:
						return v.Slice(start, end).Interface(), nil
					}
				}
			}
		}
		return nil, &ErrorInvalidArguments{Function: "Slice", Arguments: args}
	},
}
