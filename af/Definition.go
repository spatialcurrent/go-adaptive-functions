// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"reflect"
)

// Definition is a struct containing a function definition
type Definition struct {
	Inputs []interface{} // an array of types or kinds
	Output interface{}   // the type or kind returned by the defined function
}

// IsValid returns true if the args match the kinds or types or the definition.
func (d Definition) IsValid(args []interface{}) bool {
	if len(args) != len(d.Inputs) {
		return false
	}

	for i, a := range args {
		switch x := d.Inputs[i].(type) {
		case reflect.Type:
			xv := reflect.ValueOf(x)
			av := reflect.ValueOf(a)
			if !xv.IsNil() {
				if !av.IsValid() {
					return false
				}
				if ak := reflect.TypeOf(a).Kind(); ak != reflect.String && ak != reflect.Array && ak != reflect.Slice && ak != reflect.Struct && av.IsNil() {
					return false
				}
				if !av.Type().AssignableTo(x) {
					return false
				}
			}
		case reflect.Kind:
			if xv := reflect.ValueOf(x); xv.IsValid() && reflect.TypeOf(a).Kind() != x {
				return false
			}
		}
	}

	return true
}
