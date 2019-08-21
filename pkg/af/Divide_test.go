// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDivideIntsDivisible(t *testing.T) {
	out, err := Divide.ValidateRun(4, 2)
	assert.NoError(t, err)
	assert.Equal(t, 2, out)
}

func TestDivideIntsNotDivisible(t *testing.T) {
	out, err := Divide.ValidateRun(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, 0.5, out)
}

func TestDivideFloats(t *testing.T) {
	out, err := Divide.ValidateRun(4, 2.0)
	assert.NoError(t, err)
	assert.Equal(t, 4.0/2.0, out)
}

func TestDivideDurations(t *testing.T) {
	out, err := Divide.ValidateRun(time.Hour*3, time.Hour*2)
	assert.NoError(t, err)
	assert.Equal(t, 1.5, out)
}

func TestDivideDurationInt(t *testing.T) {
	out, err := Divide.ValidateRun(time.Hour*4, 2)
	assert.NoError(t, err)
	assert.Equal(t, time.Hour*2, out)
}
