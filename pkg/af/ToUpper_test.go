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

func TestToUpperString(t *testing.T) {
	out, err := ToUpper.ValidateRun("HELLO world")
	assert.NoError(t, err)
	assert.Equal(t, "HELLO WORLD", out)
}

func TestToUpperBytes(t *testing.T) {
	out, err := ToUpper.ValidateRun([]byte("HELLO world"))
	assert.NoError(t, err)
	assert.Equal(t, []byte("HELLO WORLD"), out)
}
