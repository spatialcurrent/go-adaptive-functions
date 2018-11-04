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

func intArray(args []interface{}) (interface{}, error) {

	if len(args) != 1 {
		return nil, &ErrorInvalidArguments{Function: "IntArray", Arguments: args}
	}

	t := reflect.TypeOf(args[0])
	if !(t.Kind() == reflect.Array || t.Kind() == reflect.Slice) {
		return nil, &ErrorInvalidArguments{Function: "IntArray", Arguments: args}
	}

	v := reflect.ValueOf(args[0])
	l := v.Len()
	arr := make([]int, 0, l)
	for i := 0; i < l; i++ {
		x := v.Index(i).Interface()
		xt := reflect.TypeOf(x)
		if !xt.ConvertibleTo(intType) {
			fmt.Println("xt:", xt)
			return nil, &ErrorInvalidArguments{Function: "IntArray", Arguments: args}
		}
		arr = append(arr, reflect.ValueOf(x).Convert(intType).Interface().(int))
	}

	return arr, nil
}

var ToIntArray = Function{
	Name:    "IntArray",
	Aliases: []string{"intArray"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.Array}, Output: intArrayType},
		Definition{Inputs: []interface{}{reflect.Slice}, Output: intArrayType},
	},
	Function: intArray,
}
