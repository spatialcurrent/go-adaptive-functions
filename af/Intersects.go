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

func intersects(args []interface{}) (interface{}, error) {
	av := reflect.ValueOf(args[0])
	bv := reflect.ValueOf(args[1])
	for _, v := range bv.MapKeys() {
		if av.MapIndex(v).IsValid() {
			return true, nil
		}
	}
	return false, nil
}

var Intersects = Function{
	Name:    "intersects",
	Aliases: []string{"intersects"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.Map, reflect.Map}, Output: boolType},
	},
	Function: intersects,
}
