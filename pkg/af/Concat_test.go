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

func TestConcat(t *testing.T) {
	in := []interface{}{
		[]interface{}{
			"a",
			"b",
			"c",
			[]interface{}{
				"d",
				"e",
				"f",
			},
			[]interface{}{
				"g",
				"h",
				"i",
			},
			"j",
			"k",
		},
		"l",
		"m",
		"n",
	}
	str, err := Concat.ValidateRun(in)
	assert.Nil(t, err)
	assert.Equal(t, "abcdefghijklmn", str)
}
