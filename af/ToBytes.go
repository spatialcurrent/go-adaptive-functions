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

func toBytes(args []interface{}) (interface{}, error) {
	switch x := args[0].(type) {
	case string:
		return []byte(x), nil
	case byte:
		return []byte{x}, nil
	case []byte:
		return x, nil
	}
	return nil, errors.New("invalid arguments for toBytes")
}

var ToBytes = Function{
	Name:    "ToBytes",
	Aliases: []string{"bytes"},
	Definitions: []Definition{
		Definition{
			Inputs: []interface{}{byteType},
			Output: byteArrayType,
		},
		Definition{
			Inputs: []interface{}{byteArrayType},
			Output: byteArrayType,
		},
		Definition{
			Inputs: []interface{}{stringType},
			Output: byteArrayType,
		},
	},
	Function: toBytes,
}
