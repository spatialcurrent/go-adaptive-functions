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

func stringArray(args ...interface{}) (interface{}, error) {

	if len(args) != 1 {
		return nil, &ErrInvalidArguments{Function: "StringArray", Arguments: args}
	}

	t := reflect.TypeOf(args[0])
	if !(t.Kind() == reflect.Array || t.Kind() == reflect.Slice) {
		if reflect.ValueOf(args[0]).Len() == 0 {
			return nil, &ErrInvalidArguments{Function: "StringArray", Arguments: args}
		}
	}

	v := reflect.ValueOf(args[0])
	l := v.Len()
	arr := make([]string, 0, l)
	for i := 0; i < l; i++ {
		arr = append(arr, fmt.Sprint(v.Index(i).Interface()))
	}

	return arr, nil
}

var ToStringArray = Function{
	Name:    "StringArray",
	Aliases: []string{"stringArray"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.Array}, Output: stringSliceType},
		Definition{Inputs: []interface{}{reflect.Slice}, Output: stringSliceType},
	},
	Function: stringArray,
}
