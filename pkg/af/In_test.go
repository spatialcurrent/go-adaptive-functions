// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInString(t *testing.T) {
	out, err := In.ValidateRun("hello", "hello world")
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}

func TestInMap(t *testing.T) {
	out, err := In.ValidateRun("a", map[string]interface{}{"a": "x"})
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}

func TestInSliceStrings(t *testing.T) {
	out, err := In.ValidateRun("a", []string{"a", "b"})
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}

func TestInSliceInterfaces(t *testing.T) {
	out, err := In.ValidateRun("a", []interface{}{"a", "b"})
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}

func TestInSliceStringsSliceStrings(t *testing.T) {
	out, err := In.ValidateRun([]string{"my", "name"}, []string{"hello", "world", "my", "name", "is"})
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}

func TestInIpRange(t *testing.T) {
	addr := net.ParseIP("192.0.2.100")
	n := net.IPNet{IP: net.ParseIP("192.0.2.0"), Mask: net.CIDRMask(24, 32)}

	out, err := In.ValidateRun(addr, n)
	assert.NoError(t, err)
	assert.Equal(t, true, out)

	out, err = In.ValidateRun(addr, &n)
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}
