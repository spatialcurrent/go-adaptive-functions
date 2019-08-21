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

func TestReplaceStrings(t *testing.T) {
	out, err := Replace.ValidateRun("hello world", "world", "earth")
	assert.NoError(t, err)
	assert.Equal(t, "hello earth", out)
}

func TestReplaceByte(t *testing.T) {
	out, err := Replace.ValidateRun([]byte("hello world"), []byte(" ")[0], []byte("+")[0])
	assert.NoError(t, err)
	assert.Equal(t, []byte("hello+world"), out)
}

func TestReplaceBytes(t *testing.T) {
	out, err := Replace.ValidateRun([]byte("hello world"), []byte("world"), []byte("earth"))
	assert.NoError(t, err)
	assert.Equal(t, []byte("hello earth"), out)
}
