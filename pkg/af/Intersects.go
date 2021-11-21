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

func intersects(args ...interface{}) (interface{}, error) {

	if len(args) != 2 {
		return nil, &ErrInvalidArguments{Function: "Intersects", Arguments: args}
	}

	at := reflect.TypeOf(args[0])
	bt := reflect.TypeOf(args[0])
	if at.Kind() == reflect.Map && bt.Kind() == reflect.Map {
		av := reflect.ValueOf(args[0])
		bv := reflect.ValueOf(args[1])
		for _, v := range bv.MapKeys() {
			if av.MapIndex(v).IsValid() {
				return true, nil
			}
		}
		return false, nil
	}

	if av, ok := args[0].([]float64); ok && len(av) == 4 {
		if bv, ok := args[1].([]float64); ok && len(bv) == 4 {
			return av[0] < bv[2] && av[2] > bv[0] && av[3] > bv[1] && av[1] < bv[3], nil
		}
	}

	return nil, &ErrInvalidArguments{Function: "Intersects", Arguments: args}
}

var Intersects = Function{
	Name:    "Intersects",
	Aliases: []string{"intersects"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.Map, reflect.Map}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{float64ArrayType, float64ArrayType}, Output: reflect.Bool},
	},
	Function: intersects,
}
