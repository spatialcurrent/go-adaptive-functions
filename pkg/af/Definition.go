// =================================================================
//
// Copyright (C) 2021 Spatial Current, Inc. - All Rights Reserved
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

// IsBoolean returns true if the definition returns a boolean value.
func (d Definition) IsBoolean() bool {
	if k, ok := d.Output.(reflect.Kind); ok {
		if k == reflect.Bool {
			return true
		}
	}
	return false
}

// IsInteger returns true if the definition returns an integer value.
func (d Definition) IsInteger() bool {
	if k, ok := d.Output.(reflect.Kind); ok {
		if k == reflect.Int || k == reflect.Int8 || k == reflect.Int16 || k == reflect.Int32 || k == reflect.Int64 || k == reflect.Uint || k == reflect.Uint8 || k == reflect.Uint16 || k == reflect.Uint32 || k == reflect.Uint64 {
			return true
		}
	}
	return false
}

// IsFloat returns true if the definition returns a floating point number value.
func (d Definition) IsFloat() bool {
	if k, ok := d.Output.(reflect.Kind); ok {
		if k == reflect.Float32 || k == reflect.Float64 {
			return true
		}
	}
	return false
}

// IsNumber returns true if the definition returns a numeric value.
func (d Definition) IsNumber() bool {
	return d.IsInteger() || d.IsFloat()
}

// IsString returns true if the definition returns a string value.
func (d Definition) IsString() bool {
	if k, ok := d.Output.(reflect.Kind); ok {
		if k == reflect.String {
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
