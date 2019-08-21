// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

func join(args ...interface{}) (interface{}, error) {

	if len(args) != 2 {
		return nil, &ErrInvalidArguments{Function: "Join", Arguments: args}
	}

	if b, ok := args[1].(string); ok {
		if a, ok := args[0].([]string); ok {
			return strings.Join(a, b), nil
		}
		if a, ok := args[0].(map[string]struct{}); ok {
			keys := make([]string, 0, len(a))
			for key := range a {
				keys = append(keys, key)
			}
			return strings.Join(keys, b), nil
		}
		av := reflect.ValueOf(args[0])
		at := av.Type()
		if at.Kind() == reflect.Array || at.Kind() == reflect.Slice {
			values := make([]string, 0, av.Len())
			for i := 0; i < av.Len(); i++ {
				values = append(values, fmt.Sprint(av.Index(i).Interface()))
			}
			return strings.Join(values, b), nil
		}
	} else if b, ok := args[1].([]byte); ok {
		if a, ok := args[0].([][]byte); ok {
			return bytes.Join(a, b), nil
		}
	} else if b, ok := args[1].(byte); ok {
		if a, ok := args[0].([][]byte); ok {
			return bytes.Join(a, []byte{b}), nil
		}
	}

	return nil, &ErrInvalidArguments{Function: "Join", Arguments: args}
}

var Join = Function{
	Name:    "Join",
	Aliases: []string{"join"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{uint82DSliceType, uint8Type}, Output: reflect.String},
		Definition{Inputs: []interface{}{uint82DSliceType, uint8SliceType}, Output: reflect.String},
		Definition{Inputs: []interface{}{reflect.Array, reflect.String}, Output: reflect.String},
		Definition{Inputs: []interface{}{reflect.Slice, reflect.String}, Output: reflect.String},
		Definition{Inputs: []interface{}{reflect.String, reflect.String}, Output: reflect.String},
		Definition{Inputs: []interface{}{stringSetType, reflect.String}, Output: reflect.String},
	},
	Function: join,
}
