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

func TestLookupMapStringInterface(t *testing.T) {
	out, err := Lookup.ValidateRun(map[string]interface{}{"x": "y"}, "x")
	assert.NoError(t, err)
	assert.Equal(t, "y", out)
}

func TestLookupMapStringInterfaceFallback(t *testing.T) {
	out, err := Lookup.ValidateRun(map[string]interface{}{"x": "y"}, "z", "z")
	assert.NoError(t, err)
	assert.Equal(t, "z", out)
}

func TestLookupMapIntInterface(t *testing.T) {
	out, err := Lookup.ValidateRun(map[int]interface{}{4: "y"}, 4)
	assert.NoError(t, err)
	assert.Equal(t, "y", out)
}

func TestLookupMapIntInterfaceFallback(t *testing.T) {
	out, err := Lookup.ValidateRun(map[int]interface{}{4: "y"}, 5, 6)
	assert.NoError(t, err)
	assert.Equal(t, 6, out)
}

func TestLookupStringInterface(t *testing.T) {
	out, err := Lookup.ValidateRun([]string{"x", "y"}, 1)
	assert.NoError(t, err)
	assert.Equal(t, "y", out)
}

func TestLookupStringInterfaceOverflow(t *testing.T) {
	out, err := Lookup.ValidateRun([]string{"x", "y"}, 2)
	assert.NotNil(t, err)
	assert.IsType(t, &ErrOutOfRange{}, err)
	assert.Nil(t, out)
}
