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

func TestILikeString(t *testing.T) {
	out, err := ILike.ValidateRun("hello world", "HELLO%")
	assert.NoError(t, err)
	assert.Equal(t, true, out)

	out, err = ILike.ValidateRun("hello world", "%WORLD")
	assert.NoError(t, err)
	assert.Equal(t, true, out)

	out, err = ILike.ValidateRun("hello world", "HELLO WORLD")
	assert.NoError(t, err)
	assert.Equal(t, true, out)

	out, err = ILike.ValidateRun("hello world", "HELLO _____")
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}
