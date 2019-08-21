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

func TestStringify(t *testing.T) {
	out, err := Stringify.ValidateRun(map[interface{}]interface{}{"foo": "bar"})
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"foo": "bar"}, out)
}
