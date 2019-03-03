// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"github.com/pkg/errors"
	"github.com/spatialcurrent/go-counter/counter"
)

func top(args []interface{}) (interface{}, error) {

	if len(args) == 2 {
		switch n := args[1].(type) {
		case int:
			switch c := args[0].(type) {
			case counter.Counter:
				return c.Top(n, 0), nil
			case map[string]int:
				return counter.Counter(c).Top(n, 0), nil
			}
		}
	} else if len(args) == 3 {
		switch min := args[2].(type) {
		case int:
			switch n := args[1].(type) {
			case int:
				switch c := args[0].(type) {
				case counter.Counter:
					return c.Top(n, min), nil
				case map[string]int:
					return counter.Counter(c).Top(n, min), nil
				}
			}
		}
	}

	return nil, errors.New("invalid arguments for top")

}

var Top = Function{
	Name:    "Top",
	Aliases: []string{"top"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{stringIntMapType, intType}, Output: stringArrayType},
		Definition{Inputs: []interface{}{stringIntMapType, intType, intType}, Output: stringArrayType},
	},
	Function: top,
}
