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

func TestToStringBytes(t *testing.T) {
	out, err := ToString.ValidateRun([]byte("hello world"))
	assert.NoError(t, err)
	assert.Equal(t, "hello world", out)
}

func TestToStringByte(t *testing.T) {
	out, err := ToString.ValidateRun([]byte("x")[0])
	assert.NoError(t, err)
	assert.Equal(t, "x", out)
}

func TestToStringString(t *testing.T) {
	out, err := ToString.ValidateRun("hello world")
	assert.NoError(t, err)
	assert.Equal(t, "hello world", out)
}

func TestToStringInt(t *testing.T) {
	out, err := ToString.ValidateRun(2)
	assert.NoError(t, err)
	assert.Equal(t, "2", out)
}
