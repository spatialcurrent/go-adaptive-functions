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

func TestMaxInts(t *testing.T) {
	out, err := Max.ValidateRun(2, 1, 3)
	assert.NoError(t, err)
	assert.Equal(t, 3, out)
}

func TestMaxFloats(t *testing.T) {
	out, err := Max.ValidateRun(1.11, 2.22, 3.33)
	assert.NoError(t, err)
	assert.Equal(t, 3.33, out)
}

func TestMaxVaried(t *testing.T) {
	out, err := Max.ValidateRun(1, 2.22, 3)
	assert.NoError(t, err)
	assert.Equal(t, 3, out)
}

func TestMaxComplex(t *testing.T) {
	in := []interface{}{
		[]uint8{4, 5, 6},
		[]int64{8, 9, 10},
		[]float64{0.5, 0.7},
		0.2,
	}
	out, err := Max.ValidateRun(in)
	assert.NoError(t, err)
	assert.Equal(t, int64(10), out)
}
