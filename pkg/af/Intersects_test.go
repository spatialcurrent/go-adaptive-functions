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

func TestIntersectsTrue(t *testing.T) {
	a := map[string]struct{}{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}}
	b := map[string]struct{}{"a": struct{}{}, "b": struct{}{}}
	out, err := Intersects.ValidateRun(a, b)
	assert.NoError(t, err)
	assert.Equal(t, true, out)
}

func TestIntersectsFalse(t *testing.T) {
	a := map[string]struct{}{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}}
	b := map[string]struct{}{"d": struct{}{}, "e": struct{}{}}
	out, err := Intersects.ValidateRun(a, b)
	assert.NoError(t, err)
	assert.Equal(t, false, out)
}
