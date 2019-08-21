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

func TestToBytesString(t *testing.T) {
	out, err := ToBytes.ValidateRun("abc")
	assert.NoError(t, err)
	assert.Equal(t, []byte("abc"), out)
}

func TestToBytesByte(t *testing.T) {
	out, err := ToBytes.ValidateRun([]byte("a")[0])
	assert.NoError(t, err)
	assert.Equal(t, []byte("a"), out)
}

func TestToBytesBytes(t *testing.T) {
	out, err := ToBytes.ValidateRun([]byte("abc"))
	assert.NoError(t, err)
	assert.Equal(t, []byte("abc"), out)
}
