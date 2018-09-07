// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

func repeat(args []interface{}) (interface{}, error) {

	switch value := args[0].(type) {
	case string:
		switch count := args[1].(type) {
		case int:
			return strings.Repeat(value, count), nil
		}
	case []byte:
		switch count := args[1].(type) {
		case int:
			return bytes.Repeat(value, count), nil
		}
	case byte:
		switch count := args[1].(type) {
		case int:
			return bytes.Repeat([]byte{value}, count), nil
		}
	}

	return nil, errors.New("invalid arguments " + fmt.Sprint(reflect.TypeOf(args[0])) + " , " + fmt.Sprint(reflect.TypeOf(args[1])))

}

var Repeat = Function{
	Name:    "Repeat",
	Aliases: []string{"repeat"},
	Definitions: []Definition{
		Definition{
			Inputs: []interface{}{reflect.Array, reflect.Int},
			Output: reflect.Array,
		},
		Definition{
			Inputs: []interface{}{reflect.Slice, reflect.Int},
			Output: reflect.Slice,
		},
		Definition{
			Inputs: []interface{}{reflect.String, reflect.Int},
			Output: reflect.String,
		},
		Definition{
			Inputs: []interface{}{reflect.Uint8, reflect.Int},
			Output: reflect.Slice,
		},
	},
	Function: repeat,
}
