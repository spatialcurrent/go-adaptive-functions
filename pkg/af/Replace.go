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

func replace(args ...interface{}) (interface{}, error) {

	if len(args) != 3 {
		return nil, &ErrInvalidArguments{Function: "Replace", Arguments: args}
	}

	str, ok := args[0].(string)
	if !ok {
		return nil, &ErrInvalidArguments{Function: "Replace", Arguments: args}
	}

	old, ok := args[1].(string)
	if !ok {
		return nil, &ErrInvalidArguments{Function: "Replace", Arguments: args}
	}

	new, ok := args[2].(string)
	if !ok {
		return nil, &ErrInvalidArguments{Function: "Replace", Arguments: args}
	}

	return strings.Replace(str, old, new, -1), nil
}

var Replace = Function{
	Name:    "Replace",
	Aliases: []string{"replace"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.String, reflect.String, reflect.String}, Output: reflect.String},
	},
	Function: replace,
}
