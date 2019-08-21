// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"reflect"
	"strconv"
)

func intArray(args ...interface{}) (interface{}, error) {

	if len(args) != 1 {
		return nil, &ErrInvalidArguments{Function: "IntArray", Arguments: args}
	}

	t := reflect.TypeOf(args[0])
	if !(t.Kind() == reflect.Array || t.Kind() == reflect.Slice) {
		return nil, &ErrInvalidArguments{Function: "IntArray", Arguments: args}
	}

	v := reflect.ValueOf(args[0])
	l := v.Len()
	arr := make([]int, 0, l)
	for i := 0; i < l; i++ {
		x := v.Index(i).Interface()
		if xs, ok := x.(string); ok {
			xv, err := strconv.Atoi(xs)
			if err != nil {
				return nil, &ErrInvalidArguments{Function: "IntArray", Arguments: args}
			}
			arr = append(arr, xv)
		} else {
			xt := reflect.TypeOf(x)
			if !xt.ConvertibleTo(intType) {
				return nil, &ErrInvalidArguments{Function: "IntArray", Arguments: args}
			}
			arr = append(arr, reflect.ValueOf(x).Convert(intType).Interface().(int))
		}
	}

	return arr, nil
}

var ToIntSlice = Function{
	Name:    "IntSlice",
	Aliases: []string{"intArray", "intSlice"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.Array}, Output: intSliceType},
		Definition{Inputs: []interface{}{reflect.Slice}, Output: intSliceType},
	},
	Function: intArray,
}
