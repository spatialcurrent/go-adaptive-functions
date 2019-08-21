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

func toKeys(args ...interface{}) (interface{}, error) {
	if len(args) != 1 {
		return nil, &ErrInvalidArguments{Function: "ToKeys", Arguments: args}
	}

	m := args[0]

	t := reflect.TypeOf(m)
	if t.Kind() != reflect.Map {
		return nil, &ErrInvalidArguments{Function: "ToKeys", Arguments: args}
	}

	v := reflect.ValueOf(m)

	keys := reflect.MakeSlice(reflect.SliceOf(t.Key()), 0, v.Len())
	for _, key := range v.MapKeys() {
		keys = reflect.Append(keys, key)
	}

	return keys.Interface(), nil
}

var ToKeys = Function{
	Name:    "ToKeys",
	Aliases: []string{"keys"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.Map}, Output: reflect.Slice},
	},
	Function: toKeys,
}
