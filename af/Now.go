// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"time"
)

var Now = Function{
	Name:    "Now",
	Aliases: []string{"now"},
	Definitions: []Definition{
		Definition{Inputs: make([]interface{}, 0), Output: timeType},
	},
	Function: func(args []interface{}) (interface{}, error) {
		if len(args) != 0 {
			return nil, &ErrorInvalidArguments{Function: "Now", Arguments: args}
		}
		return time.Now(), nil
	},
}
