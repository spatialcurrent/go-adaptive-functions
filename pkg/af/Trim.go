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

func trim(args ...interface{}) (interface{}, error) {

	if len(args) != 1 {
		return nil, &ErrInvalidArguments{Function: "Trim", Arguments: args}
	}

	switch a := args[0].(type) {
	case string:
		return strings.TrimSpace(a), nil
	case []string:
		b := make([]string, 0, len(a))
		for _, v := range a {
			b = append(b, strings.TrimSpace(v))
		}
		return b, nil
	case []interface{}:
		b := make([]interface{}, 0, len(a))
		for _, v := range a {
			if str, ok := v.(string); ok {
				b = append(b, strings.TrimSpace(str))
			} else {
				b = append(b, v)
			}
		}
		return b, nil
	}

	return nil, &ErrInvalidArguments{Function: "Trim", Arguments: args}
}

var Trim = Function{
	Name:    "Trim",
	Aliases: []string{"trim"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.String}, Output: reflect.String},
		Definition{Inputs: []interface{}{stringArrayType}, Output: reflect.String},
		Definition{Inputs: []interface{}{interfaceArrayType}, Output: reflect.String},
	},
	Function: trim,
}
