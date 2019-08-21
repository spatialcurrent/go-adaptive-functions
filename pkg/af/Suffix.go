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

func suffix(args ...interface{}) (interface{}, error) {

	if len(args) != 2 {
		return nil, &ErrInvalidArguments{Function: "Suffix", Arguments: args}
	}

	a := args[0]
	b := args[1]

	if av, ok := a.(string); ok {
		if bv, ok := b.(string); ok {
			return strings.HasSuffix(av, bv), nil
		}
	}

	if av, ok := a.([]byte); ok {
		if bv, ok := b.([]byte); ok {
			return bytes.HasSuffix(av, bv), nil
		}
	}

	at := reflect.TypeOf(a)
	bt := reflect.TypeOf(b)

	if at.Kind() == reflect.String && bt.Kind() == reflect.String {
		return strings.HasSuffix(a.(string), b.(string)), nil
	}

	if !(at.Kind() == reflect.Array || at.Kind() == reflect.Slice) && (bt.Kind() == reflect.Array || bt.Kind() == reflect.Slice) {
		return nil, &ErrInvalidArguments{Function: "Suffix", Arguments: args}
	}

	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	al := av.Len()
	bl := bv.Len()

	if bl > al {
		return false, nil
	}

	for i := 0; i < bl; i++ {
		if !reflect.DeepEqual(av.Index(al-i-1).Interface(), bv.Index(bl-i-1).Interface()) {
			return false, nil
		}
	}

	return true, nil

}

var Suffix = Function{
	Name:    "Suffix",
	Aliases: []string{"suffix"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.String, reflect.String}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{reflect.Array, reflect.Array}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{reflect.Array, reflect.Slice}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{reflect.Slice, reflect.Array}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{reflect.Slice, reflect.Slice}, Output: reflect.Bool},
	},
	Function: suffix,
}
