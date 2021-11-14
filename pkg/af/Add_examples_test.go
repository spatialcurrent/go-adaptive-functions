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

func ExampleAdd_ints() {
	out, err := af.Add.ValidateRun(1, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", out)
	// Output: 3
}

func ExampleAdd_floats() {
	out, err := af.Add.ValidateRun(0.1, 0.2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", out)
	// Output: 0.30000000000000004
}

func ExampleAdd_strings() {
	out, err := af.Add.ValidateRun("foo", "bar")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\v", out)
	// Output: "foobar"
}
