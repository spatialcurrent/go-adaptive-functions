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

func ilike(args ...interface{}) (interface{}, error) {

	if len(args) != 2 {
		return false, &ErrInvalidArguments{Function: "ILike", Arguments: args}
	}

	a := args[0]
	b := args[1]

	as, ok := a.(string)
	if !ok {
		return nil, &ErrInvalidArguments{Function: "ILike", Arguments: args}
	}

	bs, ok := b.(string)
	if !ok {
		return nil, &ErrInvalidArguments{Function: "ILike", Arguments: args}
	}

	re, err := regexp.Compile(strings.ToLower(
		"^" + strings.ReplaceAll(
			strings.ReplaceAll(bs, "_", "(.{1})"),
			"%",
			"(.*)",
		) + "$",
	))
	if err != nil {
		return nil, &ErrInvalidArguments{Function: "ILike", Arguments: args}
	}
	return len(re.FindString(strings.ToLower(as))) > 0, nil

}

var ILike = Function{
	Name:    "ILike",
	Aliases: []string{"ilike"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.String, reflect.String}, Output: reflect.Bool},
	},
	Function: ilike,
}
