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

func TestToItems(t *testing.T) {
	out, err := ToItems.ValidateRun(map[string]interface{}{"a": "x"})
	assert.NoError(t, err)
	assert.Equal(t, [][]interface{}{[]interface{}{"a", "x"}}, out)
}
