// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"github.com/pkg/errors"
)

func toString(args []interface{}) (interface{}, error) {
	switch x := args[0].(type) {
	case string:
		return x, nil
	case byte:
		return string([]byte{x}), nil
	case []byte:
		return string(x), nil
	}
	return nil, errors.New("invalid arguments for toString")
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
