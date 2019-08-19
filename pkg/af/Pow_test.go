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

func TestPowInts(t *testing.T) {
	out, err := Pow.ValidateRun(4, 2)
	assert.NoError(t, err)
	assert.Equal(t, 16, out)
}

func TestPowFloats(t *testing.T) {
	out, err := Pow.ValidateRun(4, 2.0)
	assert.NoError(t, err)
	assert.Equal(t, 16.0, out)
}
