// =================================================================
//
// Copyright (C) 2021 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"reflect"
	"regexp"
	"strings"
)

func like(args ...interface{}) (interface{}, error) {

	if len(args) != 2 {
		return false, &ErrInvalidArguments{Function: "Like", Arguments: args}
	}

	a := args[0]
	b := args[1]

	as, ok := a.(string)
	if !ok {
		return nil, &ErrInvalidArguments{Function: "Like", Arguments: args}
	}

	bs, ok := b.(string)
	if !ok {
		return nil, &ErrInvalidArguments{Function: "Like", Arguments: args}
	}

	re, err := regexp.Compile(
		"^" + strings.ReplaceAll(
			strings.ReplaceAll(bs, "_", "(.{1})"),
			"%",
			"(.*)",
		) + "$",
	)
	if err != nil {
		return nil, &ErrInvalidArguments{Function: "Like", Arguments: args}
	}
	return len(re.FindString(as)) > 0, nil

}

var Like = Function{
	Name:    "Like",
	Aliases: []string{"like"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.String, reflect.String}, Output: reflect.Bool},
	},
	Function: like,
}
