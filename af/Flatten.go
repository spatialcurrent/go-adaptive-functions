// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"github.com/pkg/errors"
	"reflect"
)

func flatten(args []interface{}) (interface{}, error) {

	switch arr := args[0].(type) {
	case []string:
		return arr, nil
	case []int:
		return arr, nil
	case []interface{}:
		output := make([]interface{}, 0)
		for _, x := range arr {
			if t := reflect.TypeOf(x); t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
				v := reflect.ValueOf(x)
				l := v.Len()
				for i := 0; i < l; i++ {
					output = append(output, v.Index(i).Interface())
				}
			} else {
				return nil, errors.New("invalid arguments for flatten")
			}
		}
		return output, nil
	}

	return nil, errors.New("invalid arguments for flatten")
}

var Flatten = Function{
	Name:    "Flatten",
	Aliases: []string{"flatten"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{stringArrayType}, Output: stringArrayType},
		Definition{Inputs: []interface{}{intArrayType}, Output: intArrayType},
		Definition{Inputs: []interface{}{interfaceArrayType}, Output: interfaceArrayType},
	},
	Function: flatten,
}
