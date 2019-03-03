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

func float64Array(args []interface{}) (interface{}, error) {

	if len(args) != 1 {
		return nil, &ErrorInvalidArguments{Function: "Float64Array", Arguments: args}
	}

	t := reflect.TypeOf(args[0])
	if !(t.Kind() == reflect.Array || t.Kind() == reflect.Slice) {
		if reflect.ValueOf(args[0]).Len() == 0 {
			return nil, &ErrorInvalidArguments{Function: "Float64Array", Arguments: args}
		}
	}

	v := reflect.ValueOf(args[0])
	l := v.Len()
	arr := make([]float64, 0, l)
	for i := 0; i < l; i++ {
		x := v.Index(i).Interface()
		xt := reflect.TypeOf(x)
		if !xt.ConvertibleTo(float64Type) {
			return nil, &ErrorInvalidArguments{Function: "Float64Array", Arguments: args}
		}
		arr = append(arr, reflect.ValueOf(x).Convert(float64Type).Interface().(float64))
	}

	return arr, nil
}

var ToFloat64Array = Function{
	Name:    "Float64Array",
	Aliases: []string{"float64Array"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.Array}, Output: float64ArrayType},
		Definition{Inputs: []interface{}{reflect.Slice}, Output: float64ArrayType},
	},
	Function: float64Array,
}
