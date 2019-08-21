// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

func toBytes(args ...interface{}) (interface{}, error) {
	switch x := args[0].(type) {
	case string:
		return []byte(x), nil
	case byte:
		return []byte{x}, nil
	case []byte:
		return x, nil
	}
	return nil, &ErrInvalidArguments{Function: "ToBigEndian", Arguments: args}
}

var ToBytes = Function{
	Name:    "ToBytes",
	Aliases: []string{"bytes"},
	Definitions: []Definition{
		Definition{
			Inputs: []interface{}{byteType},
			Output: byteSliceType,
		},
		Definition{
			Inputs: []interface{}{byteSliceType},
			Output: byteSliceType,
		},
		Definition{
			Inputs: []interface{}{stringType},
			Output: byteSliceType,
		},
	},
	Function: toBytes,
}
