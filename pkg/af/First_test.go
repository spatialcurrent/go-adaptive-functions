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

func TestFirstInts(t *testing.T) {
	out, err := First.ValidateRun([]interface{}{4, 2})
	assert.NoError(t, err)
	assert.Equal(t, 4, out)
}

func TestFirstFloats(t *testing.T) {
	out, err := First.ValidateRun([]interface{}{4.12, 2.22})
	assert.NoError(t, err)
	assert.Equal(t, 4.12, out)
}

func TestFirstStrings(t *testing.T) {
	out, err := First.ValidateRun([]interface{}{"a", "b"})
	assert.NoError(t, err)
	assert.Equal(t, "a", out)
}
