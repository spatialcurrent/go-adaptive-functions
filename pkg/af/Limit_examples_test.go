// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af_test

import (
	"fmt"

	"github.com/spatialcurrent/go-adaptive-functions/pkg/af"
)

// This examples shows the use of the limit function to limit the number of elements in a slice.
func Example_limit() {
	out, err := af.Limit.ValidateRun([]string{"a", "b", "c"}, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%#v", out))
	// Output: []string{"a", "b"}
}
