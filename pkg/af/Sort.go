// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"sort"
)

func sortArray(args ...interface{}) (interface{}, error) {

	descending := false
	if len(args) > 1 {
		switch args[1].(type) {
		case bool:
			descending = args[1].(bool)
		default:
			return nil, &ErrInvalidArguments{Function: "Sort", Arguments: args}
		}
	}

	switch arr := args[0].(type) {
	case []string:
		if descending {
			sort.Sort(sort.Reverse(sort.StringSlice(arr)))
		} else {
			sort.Strings(arr)
		}
		return arr, nil
	case []int:
		if descending {
			sort.Sort(sort.Reverse(sort.IntSlice(arr)))
		} else {
			sort.Ints(arr)
		}
		return arr, nil
	case []float64:
		if descending {
			sort.Sort(sort.Reverse(sort.Float64Slice(arr)))
		} else {
			sort.Float64s(arr)
		}
		return arr, nil
	}

	return nil, &ErrInvalidArguments{Function: "Sort", Arguments: args}
}

var Sort = Function{
	Name:    "Sort",
	Aliases: []string{"sort"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{stringSliceType}, Output: stringSliceType},
		Definition{Inputs: []interface{}{intArrayType}, Output: intArrayType},
		Definition{Inputs: []interface{}{stringSliceType, boolType}, Output: stringSliceType},
		Definition{Inputs: []interface{}{intArrayType, boolType}, Output: intArrayType},
	},
	Function: sortArray,
}
