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

func TestMultiplyInts(t *testing.T) {
	out, err := Multiply.ValidateRun(2, 1, 3)
	assert.NoError(t, err)
	assert.Equal(t, 6, out)
}

func TestMultiplyFloats(t *testing.T) {
	out, err := Multiply.ValidateRun(1.5, 3.0)
	assert.NoError(t, err)
	assert.Equal(t, 4.5, out)
}

func TestMultiplyVaried(t *testing.T) {
	out, err := Multiply.ValidateRun(1.5, 3)
	assert.NoError(t, err)
	assert.Equal(t, 4.5, out)
}

func TestMultiplyComplex(t *testing.T) {
	in := []interface{}{
		[]uint8{1},
		[]float64{0.5, 0.7},
		0.2,
	}
	out, err := Multiply.ValidateRun(in)
	assert.NoError(t, err)
	assert.Equal(t, 0.06999999999999999, out)
}
