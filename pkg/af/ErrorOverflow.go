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

// Error returns the error as a string.
func (e ErrOverflow) Error() string {
	return fmt.Sprintf("overflow error when trying to convert %v (%T) to %v", e.Original, e.Original, e.Target)
}
