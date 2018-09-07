// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"github.com/pkg/errors"
	"strings"
)

func toUpper(args []interface{}) (interface{}, error) {

	switch x := args[0].(type) {
	case string:
		return strings.ToUpper(x), nil
	}

	return nil, errors.New("invalid arguments for toUpper")
}

var ToUpper = Function{
	Name:    "ToUpper",
	Aliases: []string{"upper"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{stringType}, Output: stringType},
	},
	Function: toUpper,
}
