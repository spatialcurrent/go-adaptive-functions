// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"bytes"
	"reflect"
	"strings"
)

func toUpper(args ...interface{}) (interface{}, error) {

	switch x := args[0].(type) {
	case string:
		return strings.ToUpper(x), nil
	case []byte:
		return bytes.ToUpper(x), nil
	}

	return nil, &ErrInvalidArguments{Function: "ToUpper", Arguments: args}
}

var ToUpper = Function{
	Name:    "ToUpper",
	Aliases: []string{"upper"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.String}, Output: reflect.String},
		Definition{Inputs: []interface{}{byteSliceType}, Output: byteSliceType},
	},
	Function: toUpper,
}
