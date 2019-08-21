// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToLittleEndian(t *testing.T) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, int64(123))
	if err != nil {
		panic(err)
	}
	out, err := ToLittleEndian.ValidateRun(int64(123))
	assert.NoError(t, err)
	assert.Equal(t, buf.Bytes(), out)
}
