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

// IsBoolean returns true if the definition always returns a boolean value.
func (d Definition) IsBoolean(args ...interface{}) bool {
	if k, ok := d.Output.(reflect.Kind); ok {
		if k == reflect.Bool {
			return true
		}
	}
	return false
}

// IsValid returns true if the args match the kinds or types or the definition.
func (d Definition) IsValid(args ...interface{}) bool {
	if d.Inputs == nil {
		return true
	}
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
				if ak := reflect.TypeOf(a).Kind(); (ak == reflect.Chan || ak == reflect.Func || ak == reflect.Map || ak == reflect.Ptr || ak == reflect.Slice) && av.IsNil() {
					return false
				}
				if !av.Type().AssignableTo(x) {
					return false
				}
			}
		case reflect.Kind:
			xv := reflect.ValueOf(x)
			av := reflect.ValueOf(a)
			if xv.IsValid() || !xv.IsNil() {
				if !av.IsValid() {
					return false
				}
				if xv.IsValid() && av.Type().Kind() != x {
					return false
				}
			}
		}
	}

	return true
}
