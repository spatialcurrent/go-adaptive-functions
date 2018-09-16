// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"math"
)

func tileX(args []interface{}) (interface{}, error) {
	switch z := args[1].(type) {
	case int:
		switch lon := args[0].(type) {
		case int:
			return int((180.0 + float64(lon)) * (math.Pow(float64(2), float64(z)) / 360.0)), nil
		case float64:
			return int((180.0 + lon) * (math.Pow(float64(2), float64(z)) / 360.0)), nil
		}
	}
	return nil, &ErrorInvalidArguments{Function: "TileX", Arguments: args}
}

var TileX = Function{
	Name:    "TileX",
	Aliases: []string{"tileX"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{float64Type, intType}, Output: intType},
		Definition{Inputs: []interface{}{intType, intType}, Output: intType},
	},
	Function: tileX,
}
