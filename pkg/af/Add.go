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

	"github.com/spatialcurrent/go-math/pkg/math"
)

func add(args ...interface{}) (interface{}, error) {

	if len(args) != 2 {
		return nil, &ErrInvalidArguments{Function: "Add", Arguments: args}
	}

	a := args[0]
	b := args[1]

	at := reflect.TypeOf(a)
	bt := reflect.TypeOf(b)

	if at.Kind() == reflect.Map {
		if bt.Kind() != reflect.Map {
			return nil, &ErrInvalidArguments{Function: "Add", Arguments: args}
		}
		var ct reflect.Type
		var cv reflect.Value
		if bt.Key().AssignableTo(at.Key()) {
			if bt.Elem().AssignableTo(at.Elem()) {
				ct = at
				cv = reflect.MakeMap(ct)
			} else {
				ct = reflect.MapOf(at.Key(), interfaceType)
				cv = reflect.MakeMap(ct)
			}
		} else if at.Key().AssignableTo(bt.Key()) {
			if at.Elem().AssignableTo(bt.Elem()) {
				ct = bt
				cv = reflect.MakeMap(ct)
			} else {
				ct = reflect.MapOf(bt.Key(), interfaceType)
				cv = reflect.MakeMap(ct)
			}
		} else {
			ct = reflect.MapOf(interfaceType, interfaceType)
			cv = reflect.MakeMap(ct)
		}

		av := reflect.ValueOf(a)
		for _, k := range av.MapKeys() {
			cv.SetMapIndex(k.Convert(ct.Key()), av.MapIndex(k))
		}

		bv := reflect.ValueOf(b)
		for _, k := range bv.MapKeys() {
			cv.SetMapIndex(k.Convert(ct.Key()), bv.MapIndex(k))
		}
		return cv.Interface(), nil
	}

	if at.Kind() == reflect.String || bt.Kind() == reflect.String {
		return fmt.Sprint(a) + fmt.Sprint(b), nil
	}

	if at.Kind() == reflect.Array || at.Kind() == reflect.Slice {
		if !at.Elem().AssignableTo(bt.Elem()) {
			return nil, &ErrInvalidArguments{Function: "Add", Arguments: args}
		}
		av := reflect.ValueOf(a)
		bv := reflect.ValueOf(b)
		cv := reflect.MakeSlice(reflect.SliceOf(at.Elem()), 0, av.Len()+bv.Len())
		cv = reflect.AppendSlice(cv, av)
		cv = reflect.AppendSlice(cv, bv)
		return cv.Interface(), nil
	}

	return math.Add(a, b)
}

var Add = Function{
	Name:    "Add",
	Aliases: []string{"add"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.Int, reflect.Int}, Output: reflect.Int},
		Definition{Inputs: []interface{}{reflect.Int, reflect.Int64}, Output: reflect.Int},
		Definition{Inputs: []interface{}{reflect.Int, reflect.Float64}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{reflect.Int64, reflect.Int}, Output: reflect.Int},
		Definition{Inputs: []interface{}{reflect.Int64, reflect.Int64}, Output: reflect.Int},
		Definition{Inputs: []interface{}{reflect.Int64, reflect.Float64}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{reflect.Float64, reflect.Int}, Output: reflect.Int},
		Definition{Inputs: []interface{}{reflect.Float64, reflect.Int64}, Output: reflect.Int},
		Definition{Inputs: []interface{}{reflect.Float64, reflect.Float64}, Output: reflect.Float64},
		Definition{Inputs: []interface{}{reflect.String, reflect.String}, Output: reflect.String},
		Definition{Inputs: []interface{}{reflect.Int, reflect.String}, Output: reflect.String},
		Definition{Inputs: []interface{}{reflect.Int32, reflect.String}, Output: reflect.String},
		Definition{Inputs: []interface{}{reflect.Int64, reflect.String}, Output: reflect.String},
		Definition{Inputs: []interface{}{reflect.Float64, reflect.String}, Output: reflect.String},
		Definition{Inputs: []interface{}{reflect.String, reflect.Int}, Output: reflect.String},
		Definition{Inputs: []interface{}{reflect.String, reflect.Int32}, Output: reflect.String},
		Definition{Inputs: []interface{}{reflect.String, reflect.Int64}, Output: reflect.String},
		Definition{Inputs: []interface{}{reflect.String, reflect.Float64}, Output: reflect.String},
		Definition{Inputs: []interface{}{reflect.Slice, reflect.Array}, Output: reflect.Slice},
		Definition{Inputs: []interface{}{reflect.Slice, reflect.Slice}, Output: reflect.Slice},
		Definition{Inputs: []interface{}{reflect.Array, reflect.Array}, Output: reflect.Slice},
		Definition{Inputs: []interface{}{reflect.Array, reflect.Slice}, Output: reflect.Slice},
		Definition{Inputs: []interface{}{reflect.Map, reflect.Map}, Output: reflect.Map},
	},
	Function: add,
}
