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

func TestJoinInterface(t *testing.T) {
	out, err := Join.ValidateRun([]interface{}{"hello", "world", 1, 2, 3}, " ")
	assert.NoError(t, err)
	assert.Equal(t, "hello world 1 2 3", out)
}

func TestJoinString(t *testing.T) {
	out, err := Join.ValidateRun([]string{"hello", "world"}, " ")
	assert.NoError(t, err)
	assert.Equal(t, "hello world", out)
}

func TestJoinByte(t *testing.T) {
	out, err := Join.ValidateRun([][]byte{[]byte("hello"), []byte("world")}, []byte(" ")[0])
	assert.NoError(t, err)
	assert.Equal(t, []byte("hello world"), out)
}

func TestJoinByteSlice(t *testing.T) {
	out, err := Join.ValidateRun([][]byte{[]byte("hello"), []byte("world")}, []byte(" "))
	assert.NoError(t, err)
	assert.Equal(t, []byte("hello world"), out)
}
