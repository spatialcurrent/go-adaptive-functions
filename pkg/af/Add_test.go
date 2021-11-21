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

func TestAddIsBoolean(t *testing.T) {
	assert.False(t, Add.IsBoolean())
}

func TestAddInts(t *testing.T) {
	out, err := Add.ValidateRun(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, 3, out)
}

func TestAddFloats(t *testing.T) {
	out, err := Add.ValidateRun(0.1, 0.3)
	assert.NoError(t, err)
	assert.Equal(t, 0.4, out)
}

func TestAddStrings(t *testing.T) {
	out, err := Add.ValidateRun("foo", "bar")
	assert.NoError(t, err)
	assert.Equal(t, "foobar", out)
}

func TestAddSets(t *testing.T) {
	out, err := Add.ValidateRun(map[string]struct{}{"a": struct{}{}}, map[string]struct{}{"b": struct{}{}})
	assert.NoError(t, err)
	assert.Equal(t, map[string]struct{}{"a": struct{}{}, "b": struct{}{}}, out)
}

func TestAddMapsStringInterface(t *testing.T) {
	out, err := Add.ValidateRun(map[string]interface{}{"a": "x"}, map[string]interface{}{"b": 2})
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"a": "x", "b": 2}, out)
}

func TestAddMapsInterfaceInterface(t *testing.T) {
	out, err := Add.ValidateRun(map[int]interface{}{1: "x"}, map[string]interface{}{"b": 2})
	assert.NoError(t, err)
	assert.Equal(t, map[interface{}]interface{}{1: "x", "b": 2}, out)
}
