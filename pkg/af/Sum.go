// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"github.com/spatialcurrent/go-flat/pkg/flat"
	"github.com/spatialcurrent/go-math/pkg/math"
)

var Sum = Function{
	Name:        "Sum",
	Aliases:     []string{"sum"},
	Definitions: []Definition{},
	Function: func(args ...interface{}) (interface{}, error) {
		return math.Sum(flat.Flat(args))
	},
}
