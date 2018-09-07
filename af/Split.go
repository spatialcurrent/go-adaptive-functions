// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"reflect"
	"strings"
)

func split(args []interface{}) (interface{}, error) {

	t0 := reflect.TypeOf(args[0])
	t1 := reflect.TypeOf(args[0])

	if t0.Kind() == reflect.String {
		if t1.Kind() == reflect.String {
			return strings.Split(args[0].(string), args[1].(string)), nil
		}
	}

	return nil, nil

}

var Split = Function{
	Name:    "Split",
	Aliases: []string{"split"},
	Definitions: []Definition{
		Definition{
			Inputs: []interface{}{stringType, stringType},
			Output: reflect.Slice,
		},
	},
	Function: split,
}
