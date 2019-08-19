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

func toValues(args ...interface{}) (interface{}, error) {
	if len(args) != 1 {
		return nil, &ErrInvalidArguments{Function: "ToValues", Arguments: args}
	}

	m := args[0]

	t := reflect.TypeOf(m)
	if t.Kind() != reflect.Map {
		return nil, &ErrInvalidArguments{Function: "ToValues", Arguments: args}
	}

	v := reflect.ValueOf(m)

	values := reflect.MakeSlice(reflect.SliceOf(t.Elem()), 0, v.Len())
	for _, key := range v.MapKeys() {
		values = reflect.Append(values, v.MapIndex(key))
	}

	return values.Interface(), nil
}

var ToValues = Function{
	Name:    "ToValues",
	Aliases: []string{"values"},
	Definitions: []Definition{
		Definition{
			Inputs: []interface{}{reflect.Map},
			Output: interfaceArrayType,
		},
	},
	Function: toValues,
}
