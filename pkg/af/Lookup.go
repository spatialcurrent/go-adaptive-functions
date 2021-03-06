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

func lookup(args ...interface{}) (interface{}, error) {

	if len(args) != 2 && len(args) != 3 {
		return nil, &ErrInvalidArguments{Function: "Lookup", Arguments: args}
	}

	a := args[0]
	b := args[1]

	at := reflect.TypeOf(a)
	bt := reflect.TypeOf(b)

	if at.Kind() == reflect.Map {

		if !bt.ConvertibleTo(at.Key()) {
			return nil, &ErrInvalidArguments{Function: "Lookup", Arguments: args}
		}

		cv := reflect.ValueOf(a).MapIndex(reflect.ValueOf(b).Convert(at.Key()))

		if !cv.IsValid() {
			if len(args) == 3 {
				return args[2], nil
			}
			return nil, nil
		}

		return cv.Interface(), nil

	} else if at.Kind() == reflect.Array || at.Kind() == reflect.Slice {

		av := reflect.ValueOf(a)

		if bv, ok := b.(int); ok {
			if bv >= av.Len() {
				return nil, &ErrOutOfRange{Index: bv, Max: av.Len() - 1}
			}
			return reflect.ValueOf(a).Index(bv).Interface(), nil
		}

	}

	return nil, &ErrInvalidArguments{Function: "Lookup", Arguments: args}
}

var Lookup = Function{
	Name:    "Lookup",
	Aliases: []string{"lookup"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.Array, nil}, Output: nil},
		Definition{Inputs: []interface{}{reflect.Array, nil, nil}, Output: nil},
		Definition{Inputs: []interface{}{reflect.Slice, nil}, Output: nil},
		Definition{Inputs: []interface{}{reflect.Slice, nil, nil}, Output: nil},
		Definition{Inputs: []interface{}{reflect.Map, nil}, Output: nil},
		Definition{Inputs: []interface{}{reflect.Map, nil, nil}, Output: nil},
	},
	Function: lookup,
}
