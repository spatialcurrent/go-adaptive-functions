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

func TestFold(t *testing.T) {
	out, err := Fold.ValidateRun("ábc")
	assert.NoError(t, err)
	assert.Equal(t, "abc", out)

	out, err = Fold.ValidateRun("řst")
	assert.NoError(t, err)
	assert.Equal(t, "rst", out)

	out, err = Fold.ValidateRun("xýz")
	assert.NoError(t, err)
	assert.Equal(t, "xyz", out)
}
