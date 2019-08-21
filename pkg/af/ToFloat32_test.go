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

func TestToFloat32Int(t *testing.T) {
	out, err := ToFloat32.ValidateRun(12)
	assert.NoError(t, err)
	assert.Equal(t, float32(12), out)
}

func TestToFloat32String(t *testing.T) {
	out, err := ToFloat32.ValidateRun("12.132123")
	assert.NoError(t, err)
	assert.Equal(t, float32(12.132123), out)
}
