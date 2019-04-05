// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"reflect"
)

func toString(args []interface{}) (interface{}, error) {

	// If nil or the zero value return an empty string.
	if !reflect.ValueOf(args[0]).IsValid() {
		return "", nil
	}

	switch x := args[0].(type) {
	case string:
		return x, nil
	case byte:
		return string([]byte{x}), nil
	case []byte:
		return string(x), nil
	}
	return nil, &ErrorInvalidArguments{Function: "ToString", Arguments: args}
}

var ToString = Function{
	Name:    "ToString",
	Aliases: []string{"string"},
	Definitions: []Definition{
		Definition{
			Inputs: []interface{}{byteType},
			Output: stringType,
		},
		Definition{
			Inputs: []interface{}{byteArrayType},
			Output: stringType,
		},
		Definition{
			Inputs: []interface{}{stringType},
			Output: stringType,
		},
	},
	Function: toString,
}
