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

func replace(args ...interface{}) (interface{}, error) {

	if len(args) != 3 && len(args) != 4 {
		return nil, &ErrInvalidArguments{Function: "Replace", Arguments: args}
	}

	if a, ok := args[0].(string); ok {
		if b, ok := args[1].(string); ok {
			if c, ok := args[2].(string); ok {
				if len(args) == 4 {
					if d, ok := args[3].(int); ok {
						return strings.Replace(a, b, c, d), nil
					}
				} else {
					return strings.Replace(a, b, c, -1), nil
				}
			}
		}
	} else if a, ok := args[0].([]byte); ok {
		if b, ok := args[1].([]byte); ok {
			if c, ok := args[2].([]byte); ok {
				if len(args) == 4 {
					if d, ok := args[3].(int); ok {
						return bytes.Replace(a, b, c, d), nil
					}
				} else {
					return bytes.Replace(a, b, c, -1), nil
				}
			}
		} else if b, ok := args[1].(byte); ok {
			if c, ok := args[2].(byte); ok {
				if len(args) == 4 {
					if d, ok := args[3].(int); ok {
						return bytes.Replace(a, []byte{b}, []byte{c}, d), nil
					}
				} else {
					return bytes.Replace(a, []byte{b}, []byte{c}, -1), nil
				}
			}
		}
	}

	return nil, &ErrInvalidArguments{Function: "Replace", Arguments: args}
}

var Replace = Function{
	Name:    "Replace",
	Aliases: []string{"replace"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.String, reflect.String, reflect.String}, Output: reflect.String},
		Definition{Inputs: []interface{}{reflect.String, reflect.String, reflect.String, reflect.Int}, Output: reflect.String},
		Definition{Inputs: []interface{}{uint8SliceType, uint8Type, uint8Type}, Output: uint8SliceType},
		Definition{Inputs: []interface{}{uint8SliceType, uint8SliceType, uint8SliceType}, Output: uint8SliceType},
		Definition{Inputs: []interface{}{uint8SliceType, uint8SliceType, uint8SliceType, reflect.Int}, Output: uint8SliceType},
	},
	Function: replace,
}
