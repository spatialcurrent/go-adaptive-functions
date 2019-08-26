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

func TestMakeSlice(t *testing.T) {
	out, err := Make.ValidateRun("slice")
	assert.NoError(t, err)
	assert.IsType(t, []interface{}{}, out)
}

func TestMakeSliceString(t *testing.T) {
	out, err := Make.ValidateRun("slice", "string")
	assert.NoError(t, err)
	assert.IsType(t, []string{}, out)
}

func TestMakeSliceStringLength(t *testing.T) {
	out, err := Make.ValidateRun("slice", "string", 10)
	assert.NoError(t, err)
	assert.IsType(t, []string{}, out)
	assert.Len(t, out, 10)
}

func TestMakeSliceStringCapacity(t *testing.T) {
	out, err := Make.ValidateRun("slice", "string", 10, 20)
	assert.NoError(t, err)
	assert.IsType(t, []string{}, out)
	assert.Len(t, out, 10)
}

func TestMakeSet(t *testing.T) {
	out, err := Make.ValidateRun("set")
	assert.NoError(t, err)
	assert.IsType(t, map[interface{}]struct{}{}, out)
}

func TestMakeSetString(t *testing.T) {
	out, err := Make.ValidateRun("set", "string")
	assert.NoError(t, err)
	assert.IsType(t, map[string]struct{}{}, out)
}

func TestMakeSetFloat64(t *testing.T) {
	out, err := Make.ValidateRun("set", "float")
	assert.NoError(t, err)
	assert.IsType(t, map[float64]struct{}{}, out)
}

func TestMakeMap(t *testing.T) {
	out, err := Make.ValidateRun("map")
	assert.NoError(t, err)
	assert.IsType(t, map[interface{}]interface{}{}, out)
}

func TestMakeMapString(t *testing.T) {
	out, err := Make.ValidateRun("map", "string")
	assert.NoError(t, err)
	assert.IsType(t, map[string]interface{}{}, out)
}

func TestMakeMapFloat64(t *testing.T) {
	out, err := Make.ValidateRun("map", "float")
	assert.NoError(t, err)
	assert.IsType(t, map[float64]interface{}{}, out)
}

func TestMakeMapStringString(t *testing.T) {
	out, err := Make.ValidateRun("map", "string", "string")
	assert.NoError(t, err)
	assert.IsType(t, map[string]string{}, out)
}

func TestMakeMapStringStringCapacity(t *testing.T) {
	out, err := Make.ValidateRun("map", "string", "string", 10)
	assert.NoError(t, err)
	assert.IsType(t, map[string]string{}, out)
}
