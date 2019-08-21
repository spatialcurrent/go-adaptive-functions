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

func TestToFloat64SliceInts(t *testing.T) {
	out, err := ToFloat64Slice.ValidateRun([]int{12, 10})
	assert.NoError(t, err)
	assert.Equal(t, []float64{12, 10}, out)
}

func TestToFloat64SliceUint8s(t *testing.T) {
	out, err := ToFloat64Slice.ValidateRun([]uint8{uint8(12), uint8(10)})
	assert.NoError(t, err)
	assert.Equal(t, []float64{12, 10}, out)
}

func TestToFloat64SliceFloat64s(t *testing.T) {
	out, err := ToFloat64Slice.ValidateRun([]float64{12.0, 10.0})
	assert.NoError(t, err)
	assert.Equal(t, []float64{12, 10}, out)
}

func TestToFloat64SliceStrings(t *testing.T) {
	out, err := ToFloat64Slice.ValidateRun([]string{"12.0", "10.0"})
	assert.NoError(t, err)
	assert.Equal(t, []float64{12, 10}, out)
}
