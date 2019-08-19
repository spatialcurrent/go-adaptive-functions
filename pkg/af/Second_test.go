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

func TestSecondInts(t *testing.T) {
	out, err := Second.ValidateRun([]interface{}{4, 2})
	assert.NoError(t, err)
	assert.Equal(t, 2, out)
}

func TestSecondFloats(t *testing.T) {
	out, err := Second.ValidateRun([]interface{}{4.12, 2.22})
	assert.NoError(t, err)
	assert.Equal(t, 2.22, out)
}

func TestSecondStrings(t *testing.T) {
	out, err := Second.ValidateRun([]interface{}{"a", "b"})
	assert.NoError(t, err)
	assert.Equal(t, "b", out)
}
