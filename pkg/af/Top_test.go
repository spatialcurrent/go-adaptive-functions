// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"testing"

	"github.com/spatialcurrent/go-counter/pkg/counter"
	"github.com/stretchr/testify/assert"
)

func TestTop(t *testing.T) {
	c := counter.New()
	for i := 0; i < 8; i++ {
		c.Increment("x")
	}
	for i := 0; i < 4; i++ {
		c.Increment("y")
	}
	for i := 0; i < 2; i++ {
		c.Increment("z")
	}
	out, err := Top.ValidateRun(c, 1)
	assert.NoError(t, err)
	assert.Equal(t, []string{"x"}, out)
}

func TestTopNumber(t *testing.T) {
	c := counter.New()
	for i := 0; i < 8; i++ {
		c.Increment("x")
	}
	for i := 0; i < 4; i++ {
		c.Increment("y")
	}
	for i := 0; i < 2; i++ {
		c.Increment("z")
	}
	out, err := Top.ValidateRun(c, 2)
	assert.NoError(t, err)
	assert.Equal(t, []string{"x", "y"}, out)
}

func TestTopNumberMax(t *testing.T) {
	c := counter.New()
	for i := 0; i < 8; i++ {
		c.Increment("x")
	}
	for i := 0; i < 4; i++ {
		c.Increment("y")
	}
	for i := 0; i < 2; i++ {
		c.Increment("z")
	}
	out, err := Top.ValidateRun(c, 2, 5)
	assert.NoError(t, err)
	assert.Equal(t, []string{"x"}, out)
}
