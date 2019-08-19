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

func TestSplit(t *testing.T) {
	out, err := Split.ValidateRun("foo,bar", ",")
	assert.NoError(t, err)
	assert.Equal(t, []string{"foo", "bar"}, out)
}
