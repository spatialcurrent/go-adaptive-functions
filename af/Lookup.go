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

func lookup(args []interface{}) (interface{}, error) {

	if len(args) != 2 {
		return nil, &ErrorInvalidArguments{Function: "Lookup", Arguments: args}
	}

	a := args[0]
	b := args[1]

	at := reflect.TypeOf(a)
	bt := reflect.TypeOf(b)

	if at.Kind() == reflect.Map {

		if !bt.ConvertibleTo(at.Key()) {
			return nil, &ErrorInvalidArguments{Function: "Lookup", Arguments: args}
		}

		cv := reflect.ValueOf(a).MapIndex(reflect.ValueOf(b).Convert(at.Key()))
		if !cv.IsValid() {
			return nil, nil
		}
		return cv.Interface(), nil

	} else if at.Kind() == reflect.Array || at.Kind() == reflect.Slice {

		if bt.Kind() != reflect.Int {
			return nil, &ErrorInvalidArguments{Function: "Lookup", Arguments: args}
		}

		av := reflect.ValueOf(a)
		bv := b.(int)

		if bv >= av.Len() {
			return nil, &ErrorInvalidArguments{Function: "Lookup", Arguments: args}
		}

		return reflect.ValueOf(a).Index(bv).Interface(), nil

	}

	return nil, &ErrorInvalidArguments{Function: "Lookup", Arguments: args}

}

var Lookup = Function{
	Name:    "Lookup",
	Aliases: []string{"lookup"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.Array, nil}, Output: nil},
		Definition{Inputs: []interface{}{reflect.Slice, nil}, Output: nil},
		Definition{Inputs: []interface{}{reflect.Map, nil}, Output: nil},
	},
	Function: lookup,
}
