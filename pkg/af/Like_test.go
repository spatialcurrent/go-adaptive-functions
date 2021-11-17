// =================================================================
//
// Copyright (C) 2021 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLikeString(t *testing.T) {
	out, err := Like.ValidateRun("hello world", "hello%")
	assert.NoError(t, err)
	assert.Equal(t, true, out)

	out, err = Like.ValidateRun("hello world", "%world")
	assert.NoError(t, err)
	assert.Equal(t, true, out)

	out, err = Like.ValidateRun("hello world", "hello world")
	assert.NoError(t, err)
	assert.Equal(t, true, out)

	out, err = Like.ValidateRun("hello world", "hello _____")
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}
