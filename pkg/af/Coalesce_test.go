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

func TestCoalesceInts(t *testing.T) {
	out, err := Coalesce.ValidateRun([]interface{}{1, 2})
	assert.NoError(t, err)
	assert.Equal(t, 1, out)
}

func TestCoalesceStrings(t *testing.T) {
	out, err := Coalesce.ValidateRun([]string{"", "b"})
	assert.NoError(t, err)
	assert.Equal(t, "b", out)
}

func TestCoalesceNil(t *testing.T) {
	out, err := Coalesce.ValidateRun([]interface{}{nil, "b"})
	assert.NoError(t, err)
	assert.Equal(t, "b", out)
}

func TestCoalesceBool(t *testing.T) {
	out, err := Coalesce.ValidateRun([]interface{}{false, "b"})
	assert.NoError(t, err)
	assert.Equal(t, false, out)
}
