// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"reflect"
)

type Definition struct {
	Inputs []interface{}
	Output interface{}
}

func (d Definition) IsValid(args []interface{}) bool {
	if len(args) != len(d.Inputs) {
		return false
	}

	for i, a := range args {
		switch x := d.Inputs[i].(type) {
		case reflect.Type:
			if !reflect.TypeOf(a).AssignableTo(x) {
				return false
			}
		case reflect.Kind:
			if reflect.TypeOf(a).Kind() != x {
				return false
			}
		}
	}

	return true
}
