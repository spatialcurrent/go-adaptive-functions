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

func TestPrefixStrings(t *testing.T) {
	out, err := Prefix.ValidateRun("foobar", "foo")
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}

func TestPrefixBytes(t *testing.T) {
	out, err := Prefix.ValidateRun([]byte("foobar"), []byte("foo"))
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}

func TestPrefixInts(t *testing.T) {
	out, err := Prefix.ValidateRun([]int{1, 2, 3, 4}, []int{1, 2})
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}

func TestPrefixIntsOverflow(t *testing.T) {
	out, err := Prefix.ValidateRun([]int{1, 2}, []int{1, 2, 3, 4})
	assert.NoError(t, err)
	assert.Equal(t, false, out)
}

func TestPrefixArrays(t *testing.T) {
	out, err := Prefix.ValidateRun([6]int{1, 2, 3, 4, 5, 6}, [3]int{1, 2, 3})
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}

func TestPrefixSlices(t *testing.T) {
	out, err := Prefix.ValidateRun([]string{"foo", "bar"}, []string{"foo"})
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}
