// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"bytes"
	"encoding/binary"
)

func toBigEndian(args ...interface{}) (interface{}, error) {
	switch x := args[0].(type) {
	case int:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.BigEndian, int64(x))
		return buf.Bytes(), err
	case int16:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.BigEndian, x)
		return buf.Bytes(), err
	case int32:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.BigEndian, x)
		return buf.Bytes(), err
	case int64:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.BigEndian, x)
		return buf.Bytes(), err
	}
	return nil, &ErrInvalidArguments{Function: "ToBigEndian", Arguments: args}
}

var ToBigEndian = Function{
	Name:    "ToBigEndian",
	Aliases: []string{"big"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{intType}, Output: byteSliceType},
		Definition{Inputs: []interface{}{int8Type}, Output: byteSliceType},
		Definition{Inputs: []interface{}{int16Type}, Output: byteSliceType},
		Definition{Inputs: []interface{}{int32Type}, Output: byteSliceType},
		Definition{Inputs: []interface{}{int64Type}, Output: byteSliceType},
	},
	Function: toBigEndian,
}
