// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"math"
)

func tileY(args []interface{}) (interface{}, error) {
	switch z := args[1].(type) {
	case int:
		switch lat := args[0].(type) {
		case int:
			lat_rad := float64(lat) * math.Pi / 180.0
			return int((1.0 - math.Log(math.Tan(lat_rad)+(1/math.Cos(lat_rad)))/math.Pi) / 2.0 * math.Pow(float64(2), float64(z))), nil
		case float64:
			lat_rad := lat * math.Pi / 180.0
			return int((1.0 - math.Log(math.Tan(lat_rad)+(1/math.Cos(lat_rad)))/math.Pi) / 2.0 * math.Pow(float64(2), float64(z))), nil
		}
	}
	return nil, &ErrorInvalidArguments{Function: "TileY", Arguments: args}
}

var TileY = Function{
	Name:    "TileY",
	Aliases: []string{"tileY"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{float64Type, intType}, Output: intType},
		Definition{Inputs: []interface{}{intType, intType}, Output: intType},
	},
	Function: tileY,
}
