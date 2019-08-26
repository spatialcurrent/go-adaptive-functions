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

func TestWordsBytes(t *testing.T) {
	out, err := Words.ValidateRun([]byte("Hello, world.  Beautiful day!"))
	assert.NoError(t, err)
	assert.Equal(t, []string{"Hello,", "world.", "Beautiful", "day!"}, out)
}

func TestWordsString(t *testing.T) {
	out, err := Words.ValidateRun("Hello, world.  Beautiful day!")
	assert.NoError(t, err)
	assert.Equal(t, []string{"Hello,", "world.", "Beautiful", "day!"}, out)
}
