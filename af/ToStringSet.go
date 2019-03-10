// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"fmt"
	"reflect"
)

func stringSet(args []interface{}) (interface{}, error) {

	if len(args) != 1 {
		return nil, &ErrorInvalidArguments{Function: "ToStringSet", Arguments: args}
	}

	ss := map[string]struct{}{}

	t := reflect.TypeOf(args[0])

	if t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
		v := reflect.ValueOf(args[0])
		for i := 0; i < v.Len(); i++ {
			ss[fmt.Sprint(v.Index(i).Interface())] = struct{}{}
		}
		return ss, nil
	}

	if t.Kind() == reflect.Map {
		iter := reflect.ValueOf(args[0]).MapRange()
		for iter.Next() {
			ss[fmt.Sprint(iter.Key().Interface())] = struct{}{}
		}
		return ss, nil
	}

	return nil, &ErrorInvalidArguments{Function: "ToStringSet", Arguments: args}
}

var ToStringSet = Function{
	Name:    "ToStringSet",
	Aliases: []string{"stringSet"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.Slice}, Output: reflect.Map},
		Definition{Inputs: []interface{}{reflect.Map}, Output: stringSetType},
	},
	Function: stringSet,
}
