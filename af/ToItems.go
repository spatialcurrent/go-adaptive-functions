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

	m := reflect.ValueOf(args[0])
	arr := make([][]interface{}, 0, m.Len())
	for _, k := range m.MapKeys() {
		arr = append(arr, []interface{}{k.Interface(), m.MapIndex(k).Interface()})
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
