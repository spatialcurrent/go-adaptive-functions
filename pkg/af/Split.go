// =================================================================
//
// Copyright (C) 2021 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"reflect"
	"strings"
)

var Split = Function{
	Name:    "Split",
	Aliases: []string{"split"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.String, reflect.String}, Output: reflect.Slice},
		Definition{Inputs: []interface{}{reflect.String, reflect.String, intType}, Output: reflect.Slice},
	},
	Function: func(args ...interface{}) (interface{}, error) {
		if str, ok := args[0].(string); ok {
			if delim, ok := args[1].(string); ok {
				if len(args) == 2 {
					return strings.Split(str, delim), nil
				} else if len(args) == 3 {
					if n, ok := args[2].(int); ok {
						return strings.SplitN(str, delim, n), nil
					}
				}
			}
		}
		return nil, &ErrInvalidArguments{Function: "Split", Arguments: args}
	},
}
