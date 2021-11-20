// =================================================================
//
// Copyright (C) 2021 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"path/filepath"
	"reflect"
)

var Clean = Function{
	Name:    "Clean",
	Aliases: []string{"clean"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.String}, Output: reflect.String},
	},
	Function: func(args ...interface{}) (interface{}, error) {
		if str, ok := args[0].(string); ok {
			return filepath.Clean(str), nil
		}
		return nil, &ErrInvalidArguments{Function: "Clean", Arguments: args}
	},
}
