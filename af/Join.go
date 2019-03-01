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

func join(args []interface{}) (interface{}, error) {

	if len(args) != 2 {
		return nil, &ErrorInvalidArguments{Function: "Join", Arguments: args}
	}

	switch b := args[1].(type) {
	case string:
		if a, ok := args[0].([]string); ok {
			return strings.Join(a, b), nil
		}
		if a, ok := args[0].([]interface{}); ok {
			values := make([]string, 0, len(a))
			for _, obj := range a {
				values = append(values, fmt.Sprint(obj))
			}
			return strings.Join(values, b), nil
		}
		if a, ok := args[0].(map[string]struct{}); ok {
			keys := make([]string, 0, len(a))
			for key, _ := range a {
				keys = append(keys, key)
			}
			return strings.Join(keys, b), nil
		}
	case []byte:
		if a, ok := args[0].([][]byte); ok {
			return bytes.Join(a, b), nil
		}
	case byte:
		if a, ok := args[0].([][]byte); ok {
			return bytes.Join(a, []byte{b}), nil
		}
	}

	return nil, &ErrorInvalidArguments{Function: "Join", Arguments: args}
}

var Join = Function{
	Name:    "Join",
	Aliases: []string{"join"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{stringArrayType, reflect.String}, Output: reflect.String},
		Definition{Inputs: []interface{}{interfaceArrayType, reflect.String}, Output: reflect.String},
		Definition{Inputs: []interface{}{stringSetType, reflect.String}, Output: reflect.String},
	},
	Function: join,
}
