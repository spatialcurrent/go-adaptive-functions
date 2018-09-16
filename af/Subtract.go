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

func subtract(args []interface{}) (interface{}, error) {

	if len(args) != 2 {
		return nil, &ErrorInvalidArguments{Function: "Subtract", Arguments: args}
	}

	a := args[0]
	b := args[1]

	if at := reflect.TypeOf(a); at.Kind() == reflect.Map {
		bt := reflect.TypeOf(b)
		if bt.Kind() != reflect.Map {
			return nil, &ErrorInvalidArguments{Function: "Subtract", Arguments: args}
		}
		if !bt.Key().ConvertibleTo(at.Key()) {
			return nil, &ErrorInvalidArguments{Function: "Subtract", Arguments: args}
		}
		cv := reflect.MakeMap(at)
		av := reflect.ValueOf(a)
		for _, k := range av.MapKeys() {
			cv.SetMapIndex(k, av.MapIndex(k))
		}
		bv := reflect.ValueOf(b)
		for _, k := range bv.MapKeys() {
			// SetMapIndex to delete a key expect a "zero" value aka reflect.Value{}
			// https://stackoverflow.com/questions/23368843/how-to-delete-a-key-from-a-map-in-go-golang-using-reflection
			cv.SetMapIndex(k.Convert(at.Key()), reflect.Value{})
		}
		return cv.Interface(), nil
	}

	switch av := a.(type) {
	case int:
		switch bv := b.(type) {
		case int:
			return av - bv, nil
		case int64:
			return int64(av) - bv, nil
		case float64:
			return float64(av) - bv, nil
		}
	case int64:
		switch bv := b.(type) {
		case int:
			return av - int64(bv), nil
		case int64:
			return av - bv, nil
		case float64:
			return float64(av) - bv, nil
		}
	case float64:
		switch bv := b.(type) {
		case int:
			return av - float64(bv), nil
		case int64:
			return av - float64(bv), nil
		case float64:
			return av - bv, nil
		}
	}

	return nil, &ErrorInvalidArguments{Function: "Subtract", Arguments: args}
}

var Subtract = Function{
	Name:    "Subtract",
	Aliases: []string{"subtract"},
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
		Definition{Inputs: []interface{}{reflect.Map, reflect.Map}, Output: reflect.Map},
	},
	Function: subtract,
}
