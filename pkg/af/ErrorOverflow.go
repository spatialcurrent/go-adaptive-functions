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
)

// ErrOverflow is an error returned a number can't be converted to a lower bitsize.
type ErrOverflow struct {
	Original interface{}  // the original value
	Target   reflect.Type // the target type
}

// ErrOverflow returns the error as a string.
func (e ErrOverflow) Error() string {
	return "overflow error when trying to convert " + fmt.Sprint(e.Original) + "(" + fmt.Sprint(reflect.TypeOf(e.Original)) + ")" + " to " + fmt.Sprint(e.Target)
}
