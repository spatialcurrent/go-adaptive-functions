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

func toArray(args ...interface{}) (interface{}, error) {

	if len(args) != 1 {
		return nil, &ErrInvalidArguments{Function: "ToArray", Arguments: args}
	}

	m := args[0]

	t := reflect.TypeOf(m)
	if t.Kind() != reflect.Map {
		return nil, &ErrInvalidArguments{Function: "ToArray", Arguments: args}
	}

	v := reflect.ValueOf(m)

	if t.Elem().AssignableTo(emptyStructType) {
		// if elem is an empty struct assume input is a set
		keys := reflect.MakeSlice(reflect.SliceOf(t.Key()), 0, v.Len())
		for _, key := range v.MapKeys() {
			keys = reflect.Append(keys, key)
		}
		return keys.Interface(), nil
	}

	arr := make([][]interface{}, 0, v.Len())
	for _, k := range v.MapKeys() {
		arr = append(arr, []interface{}{k.Interface(), v.MapIndex(k).Interface()})
	}
	return arr, nil
}

var ToArray = Function{
	Name:    "ToArray",
	Aliases: []string{"array"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{stringSetType}, Output: stringArrayType},
		Definition{Inputs: []interface{}{intSetType}, Output: intArrayType},
		Definition{Inputs: []interface{}{interfaceSetType}, Output: interfaceArrayType},
		Definition{Inputs: []interface{}{reflect.Map}, Output: interface2DArrayType},
	},
	Function: toArray,
}
