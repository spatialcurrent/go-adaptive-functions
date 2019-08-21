// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"github.com/spatialcurrent/go-math/pkg/math"
)

var Pow = Function{
	Name:        "Pow",
	Aliases:     []string{"pow"},
	Definitions: []Definition{},
	Function: func(args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, &ErrInvalidArguments{Function: "Pow", Arguments: args}
		}
		return math.Pow(args[0], args[1])
	},
}
