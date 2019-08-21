// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"github.com/spatialcurrent/go-counter/pkg/counter"
)

func top(args ...interface{}) (interface{}, error) {
	if len(args) == 2 {
		if n, ok := args[1].(int); ok {
			if c, ok := args[0].(counter.Counter); ok {
				return counter.Counter(c).Top(n, -1, true), nil
			}
		}
	} else if len(args) == 3 {
		if min, ok := args[2].(int); ok {
			if n, ok := args[1].(int); ok {
				if c, ok := args[0].(counter.Counter); ok {
					return counter.Counter(c).Top(n, min, true), nil
				}
			}
		}
	}
	return nil, &ErrInvalidArguments{Function: "Top", Arguments: args}
}

var Top = Function{
	Name:    "Top",
	Aliases: []string{"top"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{stringIntMapType, intType}, Output: stringSliceType},
		Definition{Inputs: []interface{}{stringIntMapType, intType, intType}, Output: stringSliceType},
	},
	Function: top,
}
