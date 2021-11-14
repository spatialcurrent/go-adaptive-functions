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

func Example_coalesce() {
	out, err := af.Coalesce.ValidateRun(nil, "", "foo")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", out)
	// Output: "foo"
}
