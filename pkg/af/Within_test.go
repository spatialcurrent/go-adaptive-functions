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

func TestWithinInside2(t *testing.T) {
	out, err := Within.ValidateRun([]float64{-80.0, 36.0}, []float64{-168.39843, 13.58192, -52.73437, 72.1818})
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}

func TestWithinOutside2(t *testing.T) {
	out, err := Within.ValidateRun([]float64{-180.0, 36.0}, []float64{-168.39843, 13.58192, -52.73437, 72.1818})
	assert.NoError(t, err)
	assert.Equal(t, false, out)
}

func TestWithinInside3(t *testing.T) {
	out, err := Within.ValidateRun([]float64{-80.0, 36.0, 10.0}, []float64{-168.39843, 13.58192, 0.0, -52.73437, 72.1818, 100.0})
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}

func TestWithinOutside3(t *testing.T) {
	out, err := Within.ValidateRun([]float64{-180.0, 36.0, 10.0}, []float64{-168.39843, 13.58192, 0.0, -52.73437, 72.1818, 8.0})
	assert.NoError(t, err)
	assert.Equal(t, false, out)
}
