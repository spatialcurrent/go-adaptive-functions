// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"bytes"
	"reflect"
	"strings"
)

func repeat(args ...interface{}) (interface{}, error) {

	if len(args) != 2 {
		return nil, &ErrInvalidArguments{Function: "Repeat", Arguments: args}
	}

	if count, ok := args[1].(int); ok {

		if v, ok := args[0].(string); ok {
			return strings.Repeat(v, count), nil
		} else if v, ok := args[0].(byte); ok {
			return bytes.Repeat([]byte{v}, count), nil
		} else if v, ok := args[0].([]byte); ok {
			return bytes.Repeat(v, count), nil
		}

		t := reflect.TypeOf(args[0])
		if !(t.Kind() == reflect.Array || t.Kind() == reflect.Slice) {
			return nil, &ErrInvalidArguments{Function: "Repeat", Arguments: args}
		}
		v := reflect.ValueOf(args[0])
		arr := reflect.MakeSlice(t, 0, count*v.Len())
		for i := 0; i < count; i++ {
			arr = reflect.AppendSlice(arr, v)
		}
		return arr.Interface(), nil

	}

	return nil, &ErrInvalidArguments{Function: "Repeat", Arguments: args}

}

var Repeat = Function{
	Name:    "Repeat",
	Aliases: []string{"repeat"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.Array, reflect.Int}, Output: reflect.Array},
		Definition{Inputs: []interface{}{reflect.Slice, reflect.Int}, Output: reflect.Slice},
		Definition{Inputs: []interface{}{reflect.String, reflect.Int}, Output: reflect.String},
		Definition{Inputs: []interface{}{reflect.Uint8, reflect.Int}, Output: reflect.Slice},
	},
	Function: repeat,
}
