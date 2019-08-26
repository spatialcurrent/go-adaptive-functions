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

func makeObject(args ...interface{}) (interface{}, error) {
	if len(args) > 0 {
		switch args[0] {
		case "array", "slice":
			if len(args) > 1 {
				length := 0
				if len(args) >= 3 {
					x, err := ToInt.ValidateRun(args[2])
					if err != nil {
						return nil, &ErrInvalidArguments{Function: "Make", Arguments: args}
					}
					length = x.(int)
				}
				capacity := 0
				if len(args) >= 4 {
					x, err := ToInt.ValidateRun(args[3])
					if err != nil {
						return nil, &ErrInvalidArguments{Function: "Make", Arguments: args}
					}
					capacity = x.(int)
				} else {
					capacity = length
				}
				if capacity < length {
					return nil, &ErrInvalidCapacity{Value: capacity, Length: length}
				}
				switch args[1] {
				case "interface":
					return make([]interface{}, length, capacity), nil
				case "string":
					return make([]string, length, capacity), nil
				case "int", "integer":
					return make([]int, 0), nil
				case "float", "float64":
					return make([]float64, length, capacity), nil
				case "float32":
					return make([]float32, length, capacity), nil
				}
			} else {
				return make([]interface{}, 0), nil
			}
		case "set":
			if len(args) == 2 {
				switch args[1] {
				case "interface":
					return make(map[interface{}]struct{}), nil
				case "string":
					return make(map[string]struct{}), nil
				case "int", "integer":
					return make(map[int]struct{}), nil
				case "float", "float64":
					return make(map[float64]struct{}), nil
				case "float32":
					return make(map[float32]struct{}), nil
				}
			} else if len(args) == 1 {
				return make(map[interface{}]struct{}), nil
			}
		case "map":
			if len(args) <= 4 {
				key := reflect.TypeOf(make([]interface{}, 0)).Elem()
				value := reflect.TypeOf(make([]interface{}, 0)).Elem()
				capacity := 0
				if len(args) >= 2 {
					switch args[1] {
					case "string":
						key = reflect.TypeOf("")
					case "int", "integer":
						key = reflect.TypeOf(0)
					case "float", "float64":
						key = reflect.TypeOf(0.0)
					case "float32":
						key = reflect.TypeOf(float32(0.0))
					case "interface":
					}
				}
				if len(args) >= 3 {
					switch args[2] {
					case "string":
						value = reflect.TypeOf("")
					case "int", "integer":
						value = reflect.TypeOf(0)
					case "float", "float64":
						value = reflect.TypeOf(0.0)
					case "float32":
						value = reflect.TypeOf(float32(0.0))
					case "interface":
					}
				}
				if len(args) == 4 {
					x, err := ToInt.ValidateRun(args[3])
					if err != nil {
						return nil, &ErrInvalidArguments{Function: "Make", Arguments: args}
					}
					capacity = x.(int)
				}
				return reflect.MakeMapWithSize(reflect.MapOf(key, value), capacity).Interface(), nil
			}
		}
	}
	return nil, &ErrInvalidArguments{Function: "Make", Arguments: args}
}

var Make = Function{
	Name:    "Make",
	Aliases: []string{"make"},
	Definitions: []Definition{
		// map/set/slice
		Definition{Inputs: []interface{}{reflect.String}, Output: nil},
		// map/set/slice of x
		Definition{Inputs: []interface{}{reflect.String, reflect.String}, Output: nil},
		// slice/set of x with len
		Definition{Inputs: []interface{}{reflect.String, reflect.String, reflect.Int}, Output: nil},
		// slice of x w/ len and cap
		Definition{Inputs: []interface{}{reflect.String, reflect.String, reflect.Int, reflect.Int}, Output: nil},
		// map x of y
		Definition{Inputs: []interface{}{reflect.String, reflect.String, reflect.String}, Output: nil},
		// map of x => y with capacity
		Definition{Inputs: []interface{}{reflect.String, reflect.String, reflect.String, reflect.Int}, Output: nil},
	},
	Function: makeObject,
}
