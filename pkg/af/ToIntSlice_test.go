// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToIntSliceInts(t *testing.T) {
	out, err := ToIntSlice.ValidateRun([]int{12, 10})
	assert.NoError(t, err)
	assert.Equal(t, []int{12, 10}, out)
}

func TestToIntSliceUint8s(t *testing.T) {
	out, err := ToIntSlice.ValidateRun([]uint8{uint8(12), uint8(10)})
	assert.NoError(t, err)
	assert.Equal(t, []int{12, 10}, out)
}

func TestToIntSliceFloat64s(t *testing.T) {
	out, err := ToIntSlice.ValidateRun([]float64{12.0, 10.0})
	assert.NoError(t, err)
	assert.Equal(t, []int{12, 10}, out)
}

func TestToIntSliceStrings(t *testing.T) {
	out, err := ToIntSlice.ValidateRun([]string{"12", "10"})
	assert.NoError(t, err)
	assert.Equal(t, []int{12, 10}, out)
}
