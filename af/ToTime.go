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

func toTime(args []interface{}) (interface{}, error) {
	switch x := args[0].(type) {
	case string:
		t, err := time.Parse(time.RFC3339Nano, x)
		if err == nil {
			return t, nil
		}
		t, err = time.Parse(time.RFC3339, x)
		if err == nil {
			return t, nil
		}
		t, err = time.Parse("2006-01-02", x)
		if err == nil {
			return t, nil
		}
		t, err = time.Parse("2006-01-02 15:04:05 UTC", x)
		if err == nil {
			return t, nil
		}
	case time.Time:
		return x, nil
	}
	return nil, &ErrorInvalidArguments{Function: "ToTime", Arguments: args}
}

var ToTime = Function{
	Name:    "ToTime",
	Aliases: []string{"time"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{stringType}, Output: timeType},
		Definition{Inputs: []interface{}{timeType}, Output: timeType},
	},
	Function: toTime,
}
