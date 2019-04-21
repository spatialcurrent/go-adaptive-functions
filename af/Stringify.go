// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	stringify "github.com/spatialcurrent/go-stringify"
)

var Stringify = Function{
	Name:    "Stringify",
	Aliases: []string{"stringify"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{nil}, Output: nil},
	},
	Function: func(args []interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, &ErrorInvalidArguments{Function: "Stringify", Arguments: args}
		}
		return stringify.StringifyMapKeys(args[0]), nil
	},
}
