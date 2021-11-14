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

func Example_toKeys() {
	out, err := af.ToKeys.ValidateRun(map[string]interface{}{"a": "x"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", out)
	// Output: []string{"a"}
}
