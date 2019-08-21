// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"github.com/spatialcurrent/go-flat/pkg/flat"
)

// Flat is a function that flattens an array of arrays.
var Flat = Function{
	Name:        "Flat",
	Aliases:     []string{"flat", "flatten"},
	Definitions: []Definition{},
	Function: func(args ...interface{}) (interface{}, error) {
		return flat.Flat(args), nil
	},
}
