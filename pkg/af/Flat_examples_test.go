// =================================================================
//
// Copyright (C) 2021 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af_test

import (
	"fmt"

	"github.com/spatialcurrent/go-adaptive-functions/pkg/af"
)

func Example_flat() {
	in := []interface{}{
		[]interface{}{"a", "b"},
		[]interface{}{"c", "d"},
		[]interface{}{"e", "f"},
	}
	out, err := af.Flat.ValidateRun(in)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", out)
	// Output: []interface {}{"a", "b", "c", "d", "e", "f"}
}
