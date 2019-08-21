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

func TestRepeatStrings(t *testing.T) {
	out, err := Repeat.ValidateRun("x", 2)
	assert.NoError(t, err)
	assert.Equal(t, "xx", out)
}

func TestRepeatBytes(t *testing.T) {
	out, err := Repeat.ValidateRun([]byte("x"), 2)
	assert.NoError(t, err)
	assert.Equal(t, []byte("xx"), out)
}

func TestRepeatInts(t *testing.T) {
	out, err := Repeat.ValidateRun([]int{1, 2, 3}, 2)
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3, 1, 2, 3}, out)
}
