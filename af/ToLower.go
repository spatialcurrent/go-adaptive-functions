// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"strings"
)

func toLower(args []interface{}) (interface{}, error) {

	switch x := args[0].(type) {
	case string:
		return strings.ToLower(x), nil
	}

	return nil, &ErrorInvalidArguments{Function: "ToLower", Arguments: args}
}

var ToLower = Function{
	Name:    "ToLower",
	Aliases: []string{"lower"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{stringType}, Output: stringType},
	},
	Function: toLower,
}
