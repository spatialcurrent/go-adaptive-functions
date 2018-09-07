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

func intersects(args []interface{}) (interface{}, error) {

	at := reflect.TypeOf(args[0])
	bt := reflect.TypeOf(args[1])

	if at.AssignableTo(stringSetType) && bt.AssignableTo(stringSetType) {
		a := args[0].(map[string]struct{})
		b := args[1].(map[string]struct{})
		for v := range b {
			if _, ok := a[v]; ok {
				return true, nil
			}
		}
		return false, nil
	} else if at.AssignableTo(intSetType) && bt.AssignableTo(intSetType) {
		a := args[0].(map[int]struct{})
		b := args[1].(map[int]struct{})
		for v := range b {
			if _, ok := a[v]; ok {
				return true, nil
			}
		}
		return false, nil
	}

	return nil, errors.New("invalid arguments for intersects")

}

var Intersects = Function{
	Name:    "intersects",
	Aliases: []string{"intersects"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{stringSetType, stringSetType}, Output: boolType},
		Definition{Inputs: []interface{}{intSetType, intSetType}, Output: boolType},
	},
	Function: intersects,
}
