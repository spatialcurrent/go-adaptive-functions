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

func TestLimit(t *testing.T) {
	out, err := Limit.ValidateRun([]string{"a", "b", "c"}, 2)
	assert.NoError(t, err)
	assert.Equal(t, []string{"a", "b"}, out)
}
