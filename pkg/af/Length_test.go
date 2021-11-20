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

func TestLengthIsNumber(t *testing.T) {
	assert.True(t, Length.IsNumber())
}

func TestLengthIsInteger(t *testing.T) {
	assert.True(t, Length.IsInteger())
}

func TestLengthIsFloat(t *testing.T) {
	assert.False(t, Length.IsFloat())
}

func TestLengthArray(t *testing.T) {
	out, err := Length.ValidateRun([3]string{"a", "b", "c"})
	assert.NoError(t, err)
	assert.Equal(t, 3, out)
}

func TestLengthMap(t *testing.T) {
	out, err := Length.ValidateRun(map[string]string{"a": "x", "b": "y", "c": "z"})
	assert.NoError(t, err)
	assert.Equal(t, 3, out)
}

func TestLengthSlice(t *testing.T) {
	out, err := Length.ValidateRun([]string{"a", "b", "c"})
	assert.NoError(t, err)
	assert.Equal(t, 3, out)
}

func TestLengthString(t *testing.T) {
	out, err := Length.ValidateRun("foobar")
	assert.NoError(t, err)
	assert.Equal(t, 6, out)
}
