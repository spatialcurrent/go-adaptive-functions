// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"path"
)

var Base = Function{
	Name:    "Base",
	Aliases: []string{"base"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{stringType}, Output: stringType},
	},
	Function: func(args ...interface{}) (interface{}, error) {
		if str, ok := args[0].(string); ok {
			return path.Base(str), nil
		}
		return nil, &ErrInvalidArguments{Function: "Base", Arguments: args}
	},
}
