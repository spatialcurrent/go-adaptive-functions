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

// This examples shows the use of the join function.
func Example_join() {
	out, err := af.Join.ValidateRun([]interface{}{"a", "b", "c"}, ",")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", out)
	// Output: "a,b,c"
}
