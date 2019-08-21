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

func TestFlat(t *testing.T) {
	in := []interface{}{
		[]interface{}{"a", "b"},
		[]interface{}{"c", "d"},
	}
	out, err := Flat.ValidateRun(in)
	assert.NoError(t, err)
	assert.Equal(t, []interface{}{"a", "b", "c", "d"}, out)
}
