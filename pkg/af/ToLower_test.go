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

func TestToLowerString(t *testing.T) {
	out, err := ToLower.ValidateRun("HELLO world")
	assert.NoError(t, err)
	assert.Equal(t, "hello world", out)
}

func TestToLowerBytes(t *testing.T) {
	out, err := ToLower.ValidateRun([]byte("HELLO world"))
	assert.NoError(t, err)
	assert.Equal(t, []byte("hello world"), out)
}
