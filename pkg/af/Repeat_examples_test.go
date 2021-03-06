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

func Example_repeat() {
	out, err := af.Repeat.ValidateRun("ciao", 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: ciaociao
}
