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

func float64Slice(args ...interface{}) (interface{}, error) {

	if len(args) != 1 {
		return nil, &ErrInvalidArguments{Function: "Float64Slice", Arguments: args}
	}

	t := reflect.TypeOf(args[0])
	if !(t.Kind() == reflect.Array || t.Kind() == reflect.Slice) {
		if reflect.ValueOf(args[0]).Len() == 0 {
			return nil, &ErrInvalidArguments{Function: "Float64Slice", Arguments: args}
		}
	}

	v := reflect.ValueOf(args[0])
	l := v.Len()
	arr := make([]float64, 0, l)
	for i := 0; i < l; i++ {
		x := v.Index(i).Interface()
		if xs, ok := x.(string); ok {
			xv, err := strconv.ParseFloat(xs, 64)
			if err != nil {
				return nil, &ErrInvalidArguments{Function: "Float64Slice", Arguments: args}
			}
			arr = append(arr, xv)
		} else {
			xt := reflect.TypeOf(x)
			if !xt.ConvertibleTo(float64Type) {
				return nil, &ErrInvalidArguments{Function: "Float64Slice", Arguments: args}
			}
			arr = append(arr, reflect.ValueOf(x).Convert(float64Type).Interface().(float64))
		}
	}

	return arr, nil
}

var ToFloat64Slice = Function{
	Name:    "Float64Slice",
	Aliases: []string{"float64Array", "float64Slice"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.Array}, Output: float64ArrayType},
		Definition{Inputs: []interface{}{reflect.Slice}, Output: float64ArrayType},
	},
	Function: float64Slice,
}
