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

func TestMinInts(t *testing.T) {
	out, err := Min.ValidateRun(2, 1, 3)
	assert.NoError(t, err)
	assert.Equal(t, 1, out)
}

func TestMinFloats(t *testing.T) {
	out, err := Min.ValidateRun(1.11, 2.22, 3.33)
	assert.NoError(t, err)
	assert.Equal(t, 1.11, out)
}

func TestMinVaried(t *testing.T) {
	out, err := Min.ValidateRun(1, 2.22, 3)
	assert.NoError(t, err)
	assert.Equal(t, 1, out)
}

func TestMinComplex(t *testing.T) {
	in := []interface{}{
		[]uint8{4, 5, 6},
		[]int64{8, 9, 10},
		[]float64{0.5, 0.7},
		0.2,
	}
	out, err := Min.ValidateRun(in)
	assert.NoError(t, err)
	assert.Equal(t, 0.2, out)
}
