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

func toTitle(args ...interface{}) (interface{}, error) {

	switch x := args[0].(type) {
	case string:
		return strings.Title(x), nil
	case []byte:
		return bytes.Title(x), nil
	}

	return nil, &ErrInvalidArguments{Function: "ToTitle", Arguments: args}
}

var ToTitle = Function{
	Name:    "ToTitle",
	Aliases: []string{"title"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{stringType}, Output: stringType},
		Definition{Inputs: []interface{}{byteArrayType}, Output: byteArrayType},
	},
	Function: toTitle,
}
