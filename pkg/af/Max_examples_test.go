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

func Example_max() {
	out, err := af.Max.ValidateRun(1, 2, 4)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: 4
}
