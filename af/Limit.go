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

func limit(args []interface{}) (interface{}, error) {
	switch n := args[1].(type) {
	case int:
		v := reflect.ValueOf(args[0])
		if n > v.Len() {
			return args[0], nil
		}
		return v.Slice(0, n).Interface(), nil
	}
	return nil, errors.New("invalid arguments for limit")
}

var Limit = Function{
	Name:    "Limit",
	Aliases: []string{"limit"},
	Definitions: []Definition{
		Definition{
			Inputs: []interface{}{reflect.Slice, reflect.Int},
			Output: reflect.Slice,
		},
	},
	Function: limit,
}
