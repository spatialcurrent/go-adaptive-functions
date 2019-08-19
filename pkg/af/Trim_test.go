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

func TestTrimString(t *testing.T) {
	out, err := Trim.ValidateRun("\thello world ")
	assert.NoError(t, err)
	assert.Equal(t, "hello world", out)
}

func TestTrimSlice(t *testing.T) {
	out, err := Trim.ValidateRun([]string{"\thello ", "\nbeautiful world"})
	assert.NoError(t, err)
	assert.Equal(t, []string{"hello world"}, out)
}
