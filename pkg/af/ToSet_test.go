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

func TestToSetSet(t *testing.T) {
	out, err := ToSet.ValidateRun(map[string]struct{}{"a": struct{}{}, "b": struct{}{}})
	assert.NoError(t, err)
	assert.Equal(t, map[string]struct{}{"a": struct{}{}, "b": struct{}{}}, out)
}

func TestToSetMap(t *testing.T) {
	out, err := ToSet.ValidateRun(map[string]interface{}{"a": 1, "b": 2})
	assert.NoError(t, err)
	assert.Equal(t, map[string]struct{}{"a": struct{}{}, "b": struct{}{}}, out)
}

func TestToSetStrings(t *testing.T) {
	out, err := ToSet.ValidateRun([]string{"a", "b"})
	assert.NoError(t, err)
	assert.Equal(t, map[string]struct{}{"a": struct{}{}, "b": struct{}{}}, out)
}

func TestToSetInts(t *testing.T) {
	out, err := ToSet.ValidateRun([]int{1, 2})
	assert.NoError(t, err)
	assert.Equal(t, map[int]struct{}{1: struct{}{}, 2: struct{}{}}, out)
}
