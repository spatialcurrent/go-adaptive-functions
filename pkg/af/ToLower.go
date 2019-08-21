// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"bytes"
	"strings"
)

func toLower(args ...interface{}) (interface{}, error) {

	switch x := args[0].(type) {
	case string:
		return strings.ToLower(x), nil
	case []byte:
		return bytes.ToLower(x), nil
	}

	return nil, &ErrInvalidArguments{Function: "ToLower", Arguments: args}
}

var ToLower = Function{
	Name:    "ToLower",
	Aliases: []string{"lower"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{stringType}, Output: stringType},
		Definition{Inputs: []interface{}{byteSliceType}, Output: byteSliceType},
	},
	Function: toLower,
}
