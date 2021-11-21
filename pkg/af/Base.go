// =================================================================
//
// Copyright (C) 2021 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"path"
	"reflect"
)

var Base = Function{
	Name:    "Base",
	Aliases: []string{"base"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.String}, Output: reflect.String},
	},
	Function: func(args ...interface{}) (interface{}, error) {
		if str, ok := args[0].(string); ok {
			return path.Base(str), nil
		}
		return nil, &ErrInvalidArguments{Function: "Base", Arguments: args}
	},
}
