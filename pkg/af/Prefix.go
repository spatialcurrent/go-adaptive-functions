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

func prefix(args ...interface{}) (interface{}, error) {

	if len(args) != 2 {
		return nil, &ErrInvalidArguments{Function: "Prefix", Arguments: args}
	}

	a := args[0]
	b := args[1]

	if av, ok := a.(string); ok {
		if bv, ok := b.(string); ok {
			return strings.HasPrefix(av, bv), nil
		}
	}

	if av, ok := a.([]byte); ok {
		if bv, ok := b.([]byte); ok {
			return bytes.HasPrefix(av, bv), nil
		}
	}

	ak := reflect.TypeOf(a).Kind()
	bk := reflect.TypeOf(b).Kind()

	if !(ak == reflect.Array || ak == reflect.Slice) && (bk == reflect.Array || bk == reflect.Slice) {
		return nil, &ErrInvalidArguments{Function: "Prefix", Arguments: args}
	}

	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	if bv.Len() > av.Len() {
		return false, nil
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
