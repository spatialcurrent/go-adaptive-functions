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

func toValues(args []interface{}) (interface{}, error) {
	m := args[0]
	v := reflect.ValueOf(m)
	values := make([]interface{}, 0, v.Len())
	for _, key := range v.MapKeys() {
		values = append(values, v.MapIndex(key).Interface())
	}
	return values, nil
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
