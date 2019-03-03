// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"reflect"
	"strings"
)

func prefix(args []interface{}) (interface{}, error) {

	if len(args) != 2 {
		return nil, &ErrorInvalidArguments{Function: "Prefix", Arguments: args}
	}

	a := args[0]
	b := args[1]

	at := reflect.TypeOf(a)
	bt := reflect.TypeOf(b)

	if at.Kind() == reflect.String && bt.Kind() == reflect.String {
		return strings.HasPrefix(a.(string), b.(string)), nil
	}

	if !(at.Kind() == reflect.Array || at.Kind() == reflect.Slice) && (bt.Kind() == reflect.Array || bt.Kind() == reflect.Slice) {
		return nil, &ErrorInvalidArguments{Function: "Prefix", Arguments: args}
	}

	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	if bv.Len() > av.Len() {
		return nil, &ErrorInvalidArguments{Function: "Prefix", Arguments: args}
	}

	bl := bv.Len()
	for i := 0; i < bl; i++ {
		if !reflect.DeepEqual(av.Index(i).Interface(), bv.Index(i).Interface()) {
			return false, nil
		}
	}

	return true, nil

}

var Prefix = Function{
	Name:    "Prefix",
	Aliases: []string{"prefix"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.String, reflect.String}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{reflect.Array, reflect.Array}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{reflect.Array, reflect.Slice}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{reflect.Slice, reflect.Array}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{reflect.Slice, reflect.Slice}, Output: reflect.Bool},
	},
	Function: prefix,
}
