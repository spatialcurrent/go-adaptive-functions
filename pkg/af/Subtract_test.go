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

func TestSubtractIsBoolean(t *testing.T) {
	assert.False(t, Subtract.IsBoolean())
}

func TestSubtractIsNumber(t *testing.T) {
	assert.False(t, Subtract.IsNumber())
}

func TestSubtractInts(t *testing.T) {
	out, err := Subtract.ValidateRun(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, -1, out)
}

func TestSubtractFloats(t *testing.T) {
	out, err := Subtract.ValidateRun(0.1, 0.3)
	assert.NoError(t, err)
	assert.Equal(t, -0.19999999999999998, out)
}

func TestSubtractMapStringString(t *testing.T) {
	out, err := Subtract.ValidateRun(map[string]string{"a": "x", "b": "y"}, map[string]string{"a": ""})
	assert.NoError(t, err)
	assert.Equal(t, map[string]string{"b": "y"}, out)
}

func TestSubtractMapStringInterface(t *testing.T) {
	out, err := Subtract.ValidateRun(map[string]interface{}{"a": "x", "b": "y"}, map[string]interface{}{"a": ""})
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"b": "y"}, out)
}

func TestSubtractSets(t *testing.T) {
	out, err := Subtract.ValidateRun(map[string]interface{}{"a": struct{}{}, "b": struct{}{}}, map[string]struct{}{"a": struct{}{}})
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"b": struct{}{}}, out)
}
