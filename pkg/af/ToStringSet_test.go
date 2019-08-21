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

func TestToStringSetSet(t *testing.T) {
	out, err := ToStringSet.ValidateRun(map[string]struct{}{"a": struct{}{}, "b": struct{}{}})
	assert.NoError(t, err)
	assert.Equal(t, map[string]struct{}{"a": struct{}{}, "b": struct{}{}}, out)
}

func TestToStringSetMap(t *testing.T) {
	out, err := ToStringSet.ValidateRun(map[string]interface{}{"a": 1, "b": 2})
	assert.NoError(t, err)
	assert.Equal(t, map[string]struct{}{"a": struct{}{}, "b": struct{}{}}, out)
}

func TestToStringSetStrings(t *testing.T) {
	out, err := ToStringSet.ValidateRun([]string{"a", "b"})
	assert.NoError(t, err)
	assert.Equal(t, map[string]struct{}{"a": struct{}{}, "b": struct{}{}}, out)
}

func TestToStringSetInts(t *testing.T) {
	out, err := ToStringSet.ValidateRun([]int{1, 2})
	assert.NoError(t, err)
	assert.Equal(t, map[string]struct{}{"1": struct{}{}, "2": struct{}{}}, out)
}
