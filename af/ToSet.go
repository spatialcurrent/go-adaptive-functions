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

func toSet(args []interface{}) (interface{}, error) {

	if len(args) != 1 {
		return nil, &ErrorInvalidArguments{Function: "ToSet", Arguments: args}
	}

	t := reflect.TypeOf(args[0])
	v := reflect.ValueOf(args[0])

	if t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
		m := reflect.MakeMap(reflect.MapOf(t.Elem(), emptyStructType))
		for i := 0; i < v.Len(); i++ {
			m.SetMapIndex(v.Index(i), reflect.ValueOf(struct{}{}))
		}
		return m.Interface(), nil
	}

	if t.Kind() == reflect.Map {
		m := reflect.MakeMap(reflect.MapOf(t.Key(), emptyStructType))
		keys := v.MapKeys()
		for i := 0; i < len(keys); i++ {
			m.SetMapIndex(keys[i], reflect.ValueOf(struct{}{}))
		}
		return m.Interface(), nil
	}

	return nil, &ErrorInvalidArguments{Function: "ToSet", Arguments: args}

}

var ToSet = Function{
	Name:    "ToSet",
	Aliases: []string{"set"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.Slice}, Output: reflect.Map},
		Definition{Inputs: []interface{}{reflect.Map}, Output: reflect.Map},
	},
	Function: toSet,
}
