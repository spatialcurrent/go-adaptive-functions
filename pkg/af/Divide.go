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

var Divide = Function{
	Name:        "Divide",
	Aliases:     []string{"divide"},
	Definitions: []Definition{},
	Function: func(args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, &ErrInvalidArguments{Function: "Divide", Arguments: args}
		}
		return math.Divide(args[0], args[1])
	},
}
