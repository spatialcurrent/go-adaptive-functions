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

func toSet(args []interface{}) (interface{}, error) {

	t := reflect.TypeOf(args[0])
	s := reflect.ValueOf(args[0])

	m := reflect.MakeMap(reflect.MapOf(t.Elem(), emptyStructType))
	for i := 0; i < s.Len(); i++ {
		m.SetMapIndex(s.Index(i), reflect.ValueOf(struct{}{}))
	}

	return m.Interface(), nil
}

var ToSet = Function{
	Name:    "ToSet",
	Aliases: []string{"set"},
	Definitions: []Definition{
		Definition{
			Inputs: []interface{}{reflect.Slice},
			Output: reflect.Map,
		},
	},
	Function: toSet,
}
