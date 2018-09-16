// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"github.com/spatialcurrent/go-counter/counter"
)

func bottom(args []interface{}) (interface{}, error) {

	if len(args) == 2 {
		switch n := args[1].(type) {
		case int:
			switch c := args[0].(type) {
			case counter.Counter:
				return c.Bottom(n, 0), nil
			case map[string]int:
				return counter.Counter(c).Bottom(n, 0), nil
			}
		}
	} else if len(args) == 3 {
		switch max := args[2].(type) {
		case int:
			switch n := args[1].(type) {
			case int:
				switch c := args[0].(type) {
				case counter.Counter:
					return c.Bottom(n, max), nil
				case map[string]int:
					return counter.Counter(c).Bottom(n, max), nil
				}
			}
		}
	}

	return nil, &ErrorInvalidArguments{Function: "Bottom", Arguments: args}

}

var Bottom = Function{
	Name:    "Bottom",
	Aliases: []string{"bottom"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{stringIntMapType, intType}, Output: stringArrayType},
		Definition{Inputs: []interface{}{stringIntMapType, intType, intType}, Output: stringArrayType},
	},
	Function: bottom,
}
