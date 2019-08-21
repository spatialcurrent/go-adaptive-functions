// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"encoding/json"
	"reflect"

	"github.com/pkg/errors"
)

var JSON = Function{
	Name:    "json",
	Aliases: []string{"json"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{nil}, Output: reflect.String},
	},
	Function: func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, &ErrInvalidArguments{Function: "Json", Arguments: args}
		}
		out, err := json.Marshal(args[0])
		if err != nil {
			return err, errors.Wrap(err, "error marshaling to json")
		}
		return out, nil
	},
}
