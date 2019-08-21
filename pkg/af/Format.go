// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func format(args ...interface{}) (interface{}, error) {
	switch x := args[0].(type) {
	case time.Time:
		if str, ok := args[1].(string); ok {
			str = strings.ReplaceAll(str, "YYYY", "2006")
			str = strings.ReplaceAll(str, "MM", "01")
			str = strings.ReplaceAll(str, "DD", "02")
			str = strings.ReplaceAll(str, "h", "3")
			str = strings.ReplaceAll(str, "hh", "15")
			str = strings.ReplaceAll(str, "mm", "04")
			return x.Format(str), nil
		}
	case string:
		if len(args) > 1 {
			return fmt.Sprintf(x, args[1:]...), nil
		}
		return x, nil
	}
	return nil, &ErrInvalidArguments{Function: "Format", Arguments: args}
}

var Format = Function{
	Name:    "Format",
	Aliases: []string{"fmt", "format"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{timeType, reflect.String}, Output: reflect.String},
		Definition{Inputs: nil, Output: reflect.String},
	},
	Function: format,
}
