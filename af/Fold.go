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
)

func fold(args []interface{}) (interface{}, error) {

	if len(args) != 1 {
		return nil, &ErrorInvalidArguments{Function: "Fold", Arguments: args}
	}

	if str, ok := args[0].(string); ok {
		var b bytes.Buffer
		for _, c := range str {
			switch c {
			case 'á':
				b.WriteRune('a')
			case 'ř':
				b.WriteRune('r')
			case 'ý':
				b.WriteRune('y')
			default:
				b.WriteRune(c)
			}
		}
		return b.String(), nil
	}

	return nil, &ErrorInvalidArguments{Function: "Fold", Arguments: args}

}

var Fold = Function{
	Name:    "Fold",
	Aliases: []string{"fold"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.String}, Output: reflect.String},
	},
	Function: fold,
}
