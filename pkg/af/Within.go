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

var Within = Function{
	Name:    "Within",
	Aliases: []string{"within"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{float64ArrayType, float64ArrayType}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{intArrayType, intArrayType}, Output: reflect.Bool},
	},
	Function: func(args ...interface{}) (interface{}, error) {
		if coordinate, ok := args[0].([]float64); ok {
			if extent, ok := args[1].([]float64); ok {
				dims := len(coordinate)
				if dims*2 == len(extent) {
					for i, x := range coordinate {
						if x < extent[i] || x > extent[i+dims] {
							return false, nil
						}
					}
					return true, nil
				}
			}
		} else if coordinate, ok := args[0].([]int); ok {
			if extent, ok := args[1].([]int); ok {
				dims := len(coordinate)
				if dims*2 == len(extent) {
					for i, x := range coordinate {
						if x < extent[i] || x > extent[i+dims] {
							return false, nil
						}
					}
					return true, nil
				}
			}
		}
		return nil, &ErrInvalidArguments{Function: "Within", Arguments: args}
	},
}
