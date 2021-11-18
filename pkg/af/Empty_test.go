// =================================================================
//
// Copyright (C) 2021 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptySlice(t *testing.T) {
	out, err := Empty.ValidateRun([]interface{}{4, 2})
	assert.NoError(t, err)
	assert.False(t, out.(bool))

	out, err = Empty.ValidateRun([]interface{}{})
	assert.NoError(t, err)
	assert.True(t, out.(bool))
}

func TestEmptyMap(t *testing.T) {
	out, err := Empty.ValidateRun(map[string]string{"foo": "bar"})
	assert.NoError(t, err)
	assert.False(t, out.(bool))

	out, err = Empty.ValidateRun(map[string]string{})
	assert.NoError(t, err)
	assert.True(t, out.(bool))
}

func TestEmptyString(t *testing.T) {
	out, err := Empty.ValidateRun("foo bar")
	assert.NoError(t, err)
	assert.False(t, out.(bool))

	out, err = Empty.ValidateRun("")
	assert.NoError(t, err)
	assert.True(t, out.(bool))
}
