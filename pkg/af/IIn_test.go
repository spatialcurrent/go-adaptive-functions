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

func TestIInIsBoolean(t *testing.T) {
	assert.True(t, IIn.IsBoolean())
}

func TestIInString(t *testing.T) {
	out, err := IIn.ValidateRun("hello", "Hello World")
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}

func TestIInMap(t *testing.T) {
	out, err := IIn.ValidateRun("A", map[string]interface{}{"a": "x"})
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}

func TestIInSliceStrings(t *testing.T) {
	out, err := IIn.ValidateRun("A", []string{"a", "b"})
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}

func TestIInSliceInterfaces(t *testing.T) {
	out, err := IIn.ValidateRun("A", []interface{}{"a", "b"})
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}
