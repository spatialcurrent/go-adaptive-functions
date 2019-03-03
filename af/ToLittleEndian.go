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

func toLittleEndian(args []interface{}) (interface{}, error) {
	switch x := args[0].(type) {
	case int:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, int64(x))
		return buf.Bytes(), err
	case int16:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, x)
		return buf.Bytes(), err
	case int32:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, x)
		return buf.Bytes(), err
	case int64:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, x)
		return buf.Bytes(), err
	}
	return nil, &ErrorInvalidArguments{Function: "ToLittleEndian", Arguments: args}
}

var ToLittleEndian = Function{
	Name:    "ToLittleEndian",
	Aliases: []string{"little"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{intType}, Output: byteArrayType},
		Definition{Inputs: []interface{}{int8Type}, Output: byteArrayType},
		Definition{Inputs: []interface{}{int16Type}, Output: byteArrayType},
		Definition{Inputs: []interface{}{int32Type}, Output: byteArrayType},
		Definition{Inputs: []interface{}{int64Type}, Output: byteArrayType},
	},
	Function: toLittleEndian,
}
