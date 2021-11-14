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

// This examples shows the use of the iin function to test if a key exists in a map.
func Example_iInMap() {
	out, err := af.IIn.ValidateRun("A", map[string]interface{}{"a": "x", "b": "y"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", out)
	// Output: true
}

// This examples shows the use of the in function to test if an element is in a slice.
func Example_iInSlice() {
	out, err := af.IIn.ValidateRun("A", []string{"a", "b", "c"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", out)
	// Output: true
}
