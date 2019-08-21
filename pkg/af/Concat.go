// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"reflect"

	"github.com/spatialcurrent/go-stringify/pkg/stringify"
)

var Concat = Function{
	Name:    "Concat",
	Aliases: []string{"concat"},
	Definitions: []Definition{
		Definition{Inputs: nil, Output: reflect.String},
	},
	Function: func(args ...interface{}) (interface{}, error) {
		return stringify.Concat(args, stringify.NewDefaultStringer())
	},
}
