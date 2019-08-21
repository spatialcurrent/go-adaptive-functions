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

func TestMeanInts(t *testing.T) {
	out, err := Mean.ValidateRun(2, 1, 3)
	assert.NoError(t, err)
	assert.Equal(t, 2.0, out)
}

func TestMeanFloats(t *testing.T) {
	out, err := Mean.ValidateRun(1.5, 3.0)
	assert.NoError(t, err)
	assert.Equal(t, 2.25, out)
}

func TestMeanVaried(t *testing.T) {
	out, err := Mean.ValidateRun(1.5, 3)
	assert.NoError(t, err)
	assert.Equal(t, 2.25, out)
}

func TestMeanComplex(t *testing.T) {
	in := []interface{}{
		[]uint8{1},
		[]float64{0.5, 0.7},
		0.2,
	}
	out, err := Mean.ValidateRun(in)
	assert.NoError(t, err)
	assert.Equal(t, 0.6000000000000001, out)
}
