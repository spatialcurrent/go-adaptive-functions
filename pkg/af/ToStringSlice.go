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

func stringSlice(args ...interface{}) (interface{}, error) {

	if len(args) != 1 {
		return nil, &ErrInvalidArguments{Function: "ToStringSlice", Arguments: args}
	}

	if str, ok := args[0].(string); ok {
		return []string{str}, nil
	}

	t := reflect.TypeOf(args[0])
	if !(t.Kind() == reflect.Array || t.Kind() == reflect.Slice) {
		if reflect.ValueOf(args[0]).Len() == 0 {
			return nil, &ErrInvalidArguments{Function: "ToStringSlice", Arguments: args}
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

var ToStringSlice = Function{
	Name:    "StringSlice",
	Aliases: []string{"stringArray", "stringSlice"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.Array}, Output: stringSliceType},
		Definition{Inputs: []interface{}{reflect.Slice}, Output: stringSliceType},
		Definition{Inputs: []interface{}{reflect.String}, Output: stringSliceType},
	},
	Function: stringSlice,
}
