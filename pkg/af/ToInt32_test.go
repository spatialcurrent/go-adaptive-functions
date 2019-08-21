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

func TestToInt32Int(t *testing.T) {
	out, err := ToInt32.ValidateRun(12)
	assert.NoError(t, err)
	assert.Equal(t, int32(12), out)
}

func TestToInt32String(t *testing.T) {
	out, err := ToInt32.ValidateRun("12")
	assert.NoError(t, err)
	assert.Equal(t, int32(12), out)
}
