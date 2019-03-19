// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"fmt"
	"reflect"
)

func concat(args []interface{}) (interface{}, error) {

	if len(args) != 1 {
		return nil, &ErrorInvalidArguments{Function: "Concat", Arguments: args}
	}

	t := reflect.TypeOf(args[0])

	if t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
		v := reflect.ValueOf(args[0])
		output := ""
		for i := 0; i < v.Len(); i++ {
			output += fmt.Sprint(v.Index(i).Interface())
		}
		return output, nil
	}

	return nil, &ErrorInvalidArguments{Function: "Concat", Arguments: args}
}

var Concat = Function{
	Name:    "Concat",
	Aliases: []string{"concat"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.Array}, Output: reflect.String},
		Definition{Inputs: []interface{}{reflect.Slice}, Output: reflect.String},
	},
	Function: concat,
}
