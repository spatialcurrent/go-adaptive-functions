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

func TestSumInts(t *testing.T) {
	out, err := Sum.ValidateRun(2, 1, 3)
	assert.NoError(t, err)
	assert.Equal(t, 6, out)
}

func TestSumFloats(t *testing.T) {
	out, err := Sum.ValidateRun(1.11, 2.22)
	assert.NoError(t, err)
	assert.Equal(t, 3.33, out)
}

func TestSumVaried(t *testing.T) {
	out, err := Sum.ValidateRun(1, 2.22)
	assert.NoError(t, err)
	assert.Equal(t, 3.22, out)
}

func TestSumComplex(t *testing.T) {
	in := []interface{}{
		[]uint8{4, 5, 6},
		[]int64{8, 9, 10},
		[]float64{0.5, 0.7},
		0.2,
	}
	out, err := Sum.ValidateRun(in)
	assert.NoError(t, err)
	assert.Equal(t, 43.400000000000006, out)
}
