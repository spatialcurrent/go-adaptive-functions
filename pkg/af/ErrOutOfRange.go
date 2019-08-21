// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"fmt"
)

// ErrOutOfRange is an error returned when an array or slice index is out of range.
type ErrOutOfRange struct {
	Index int
	Max   int
}

// Error returns the error as a string.
func (e ErrOutOfRange) Error() string {
	return fmt.Sprintf("index %d is out of range, maximum value is %d", e.Index, e.Max)
}
