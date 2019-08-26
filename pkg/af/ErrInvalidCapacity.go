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

// ErrInvalidCapacity is an error returned when the capacity for a slice is too low
// since it must always be greater than or equal to the length.
type ErrInvalidCapacity struct {
	Value  int
	Length int
}

// Error returns the error as a string.
func (e ErrInvalidCapacity) Error() string {
	return fmt.Sprintf("capacity %d is too low, must be greater than or equal to length %d", e.Value, e.Length)
}
