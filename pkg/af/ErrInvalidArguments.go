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
)

// ErrInvalidArguments is an error returned when a function is called with invalid arguments.
type ErrInvalidArguments struct {
	Function  string        // the name of the Function
	Arguments []interface{} // the arguments for the function
}

// Error returns the error as a string.
func (e ErrInvalidArguments) Error() string {
	str := "invalid arguments for function " + e.Function + " with types "
	types := make([]string, 0, len(e.Arguments))
	for _, a := range e.Arguments {
		types = append(types, fmt.Sprint(reflect.TypeOf(a)))
	}
	str += strings.Join(types, ", ")
	return str
}
