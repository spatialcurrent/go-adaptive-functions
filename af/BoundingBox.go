// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"reflect"
)

func boundingBox(args []interface{}) (interface{}, error) {

	if len(args) != 1 {
		return nil, &ErrorInvalidArguments{Function: "BoundingBox", Arguments: args}
	}

	t := reflect.TypeOf(args[0])
	if !(t.Kind() == reflect.Array || t.Kind() == reflect.Slice) {
		if reflect.ValueOf(args[0]).Len() == 0 {
			return nil, &ErrorInvalidArguments{Function: "BoundingBox", Arguments: args}
		}
	}

	v := reflect.ValueOf(args[0])
	l := v.Len()

	if l == 0 {
		return nil, &ErrorInvalidArguments{Function: "BoundingBox", Arguments: args}
	}

	a := v.Index(0).Interface()
	at := reflect.TypeOf(a)

	if !(at.Kind() == reflect.Array || at.Kind() == reflect.Slice) {
		return nil, &ErrorInvalidArguments{Function: "BoundingBox", Arguments: args}
	}

	av := reflect.ValueOf(a)
	if av.Len() < 2 {
		return nil, &ErrorInvalidArguments{Function: "BoundingBox", Arguments: args}
	}

	x := av.Index(0).Interface()
	y := av.Index(1).Interface()

	xt := reflect.TypeOf(x)
	if !xt.ConvertibleTo(float64Type) {
		return nil, &ErrorInvalidArguments{Function: "Float64Array", Arguments: args}
	}

	yt := reflect.TypeOf(y)
	if !yt.ConvertibleTo(float64Type) {
		return nil, &ErrorInvalidArguments{Function: "Float64Array", Arguments: args}
	}

	xf := reflect.ValueOf(x).Convert(float64Type).Interface().(float64)
	yf := reflect.ValueOf(y).Convert(float64Type).Interface().(float64)

	bbox := []float64{xf, yf, xf, yf}
	for i := 0; i < l; i++ {

		a := v.Index(i).Interface()
		at := reflect.TypeOf(a)

		if !(at.Kind() == reflect.Array || at.Kind() == reflect.Slice) {
			return nil, &ErrorInvalidArguments{Function: "BoundingBox", Arguments: args}
		}

		av := reflect.ValueOf(a)
		if av.Len() < 2 {
			return nil, &ErrorInvalidArguments{Function: "BoundingBox", Arguments: args}
		}

		x := av.Index(0).Interface()
		y := av.Index(1).Interface()

		xt := reflect.TypeOf(x)
		if !xt.ConvertibleTo(float64Type) {
			return nil, &ErrorInvalidArguments{Function: "Float64Array", Arguments: args}
		}

		yt := reflect.TypeOf(y)
		if !yt.ConvertibleTo(float64Type) {
			return nil, &ErrorInvalidArguments{Function: "Float64Array", Arguments: args}
		}

		xf := reflect.ValueOf(x).Convert(float64Type).Interface().(float64)
		yf := reflect.ValueOf(y).Convert(float64Type).Interface().(float64)

		if xf < bbox[0] {
			bbox[0] = xf
		}
		if yf < bbox[1] {
			bbox[1] = yf
		}
		if xf > bbox[2] {
			bbox[2] = xf
		}
		if yf > bbox[3] {
			bbox[3] = yf
		}

	}

	return bbox, nil
}

var BoundingBox = Function{
	Name:    "BoundingBox",
	Aliases: []string{"boundingBox", "bbox"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.Array}, Output: float64ArrayType},
		Definition{Inputs: []interface{}{reflect.Slice}, Output: float64ArrayType},
	},
	Function: boundingBox,
}
