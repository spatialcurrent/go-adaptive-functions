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

func Example_concat1() {
	out, err := af.Concat.ValidateRun("foo", "bar")
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%#v", out))
	// Output: "foobar"
}

func Example_concat2() {
	out, err := af.Concat.ValidateRun("1", "2", 3)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%#v", out))
	// Output: "123"
}

func Example_concat3() {
	in := []interface{}{
		[]interface{}{
			"a",
			"b",
			"c",
			[]interface{}{
				"d",
				"e",
				"f",
			},
			[]interface{}{
				"g",
				"h",
				"i",
			},
			"j",
			"k",
		},
		"l",
		"m",
		"n",
	}
	out, err := af.Concat.ValidateRun(in)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%#v", out))
	// Output: "abcdefghijklmn"
}
