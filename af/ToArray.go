// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"github.com/pkg/errors"
	"reflect"
)

func toArray(args []interface{}) (interface{}, error) {

	switch set := args[0].(type) {
	case map[string]struct{}:
		arr := make([]string, 0, len(set))
		for v := range set {
			arr = append(arr, v)
		}
		return arr, nil
	case map[int]struct{}:
		arr := make([]int, 0, len(set))
		for v := range set {
			arr = append(arr, v)
		}
		return arr, nil
	case map[interface{}]struct{}:
		arr := make([]interface{}, 0, len(set))
		for v := range set {
			arr = append(arr, v)
		}
		return arr, nil
	}

	t := reflect.TypeOf(args[0])
	if t.Kind() == reflect.Map {
		m := reflect.ValueOf(args[0])
		arr := make([][]interface{}, 0, m.Len())
		for _, k := range m.MapKeys() {
			arr = append(arr, []interface{}{k.Interface(), m.MapIndex(k).Interface()})
		}
		return arr, nil
	}

	return nil, errors.New("invalid arguments for toArray")
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
