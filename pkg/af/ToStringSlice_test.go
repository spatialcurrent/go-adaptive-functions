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

func TestToStringSliceString(t *testing.T) {
	out, err := ToStringSlice.ValidateRun("foo")
	assert.NoError(t, err)
	assert.Equal(t, []string{"foo"}, out)
}

func TestToStringSliceStrings(t *testing.T) {
	out, err := ToStringSlice.ValidateRun([]string{"foo", "bar"})
	assert.NoError(t, err)
	assert.Equal(t, []string{"foo", "bar"}, out)
}

func TestToStringSliceInts(t *testing.T) {
	out, err := ToStringSlice.ValidateRun([]int{1, 2})
	assert.NoError(t, err)
	assert.Equal(t, []string{"1", "2"}, out)
}

func TestToStringSliceVaried(t *testing.T) {
	out, err := ToStringSlice.ValidateRun([]interface{}{1, "foo", "bar"})
	assert.NoError(t, err)
	assert.Equal(t, []string{"1", "foo", "bar"}, out)
}
