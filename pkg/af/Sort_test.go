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

func TestSortInts(t *testing.T) {
	out, err := Sort.ValidateRun([]int{3, 2, 1})
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, out)
}

func TestSortIntsReversed(t *testing.T) {
	out, err := Sort.ValidateRun([]int{1, 2, 3}, true)
	assert.NoError(t, err)
	assert.Equal(t, []int{3, 2, 1}, out)
}

func TestSortStrings(t *testing.T) {
	out, err := Sort.ValidateRun([]string{"c", "b", "a"})
	assert.NoError(t, err)
	assert.Equal(t, []string{"a", "b", "c"}, out)
}

func TestSortStringsReversed(t *testing.T) {
	out, err := Sort.ValidateRun([]string{"a", "b", "c"}, true)
	assert.NoError(t, err)
	assert.Equal(t, []string{"c", "b", "a"}, out)
}
