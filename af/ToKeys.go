// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"fmt"
	"reflect"
)

func toKeys(args []interface{}) (interface{}, error) {
	m := args[0]
	v := reflect.ValueOf(m)
	keys := make([]string, 0, v.Len())
	for _, k := range v.MapKeys() {
		keys = append(keys, fmt.Sprint(k.Interface()))
	}
	return keys, nil
}

var ToKeys = Function{
	Name:    "ToKeys",
	Aliases: []string{"keys"},
	Definitions: []Definition{
		Definition{
			Inputs: []interface{}{reflect.Map},
			Output: stringArrayType,
		},
	},
	Function: toKeys,
}
