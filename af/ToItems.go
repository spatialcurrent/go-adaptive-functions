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

func toItems(args []interface{}) (interface{}, error) {
	if len(args) != 1 {
		return nil, &ErrorInvalidArguments{Function: "ToItems", Arguments: args}
	}

	m := args[0]

	t := reflect.TypeOf(m)
	if t.Kind() != reflect.Map {
		return nil, &ErrorInvalidArguments{Function: "ToItems", Arguments: args}
	}

	v := reflect.ValueOf(m)

	arr := make([][]interface{}, 0, v.Len())
	for _, k := range v.MapKeys() {
		arr = append(arr, []interface{}{k.Interface(), v.MapIndex(k).Interface()})
	}

	return arr, nil
}

var ToItems = Function{
	Name:    "ToItems",
	Aliases: []string{"items"},
	Definitions: []Definition{
		Definition{
			Inputs: []interface{}{reflect.Map},
			Output: reflect.Slice,
		},
	},
	Function: toItems,
}
