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

func TestSuffixStrings(t *testing.T) {
	out, err := Suffix.ValidateRun("foobar", "bar")
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}

func TestSuffixBytes(t *testing.T) {
	out, err := Suffix.ValidateRun([]byte("foobar"), []byte("bar"))
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}

func TestSuffixInts(t *testing.T) {
	out, err := Suffix.ValidateRun([]int{1, 2, 3, 4}, []int{3, 4})
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}

func TestSuffixArrays(t *testing.T) {
	out, err := Suffix.ValidateRun([6]int{1, 2, 3, 4, 5, 6}, [3]int{4, 5, 6})
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}

func TestSuffixSlices(t *testing.T) {
	out, err := Suffix.ValidateRun([]string{"foo", "bar"}, []string{"bar"})
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}
