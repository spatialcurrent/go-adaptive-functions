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

func TestToTitleString(t *testing.T) {
	out, err := ToUpper.ValidateRun("hello world")
	assert.NoError(t, err)
	assert.Equal(t, "Hello World", out)
}

func TestToTitleBytes(t *testing.T) {
	out, err := ToUpper.ValidateRun([]byte("hello world"))
	assert.NoError(t, err)
	assert.Equal(t, []byte("Hello World"), out)
}
