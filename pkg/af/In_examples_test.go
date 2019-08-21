// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af_test

import (
	"fmt"
	"net"

	"github.com/spatialcurrent/go-adaptive-functions/pkg/af"
)

// This examples shows the use of the in function to test if a key exists in a map.
func Example_inMap() {
	out, err := af.In.ValidateRun("a", map[string]interface{}{"a": "x", "b": "y"})
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%#v", out))
	// Output: true
}

// This examples shows the use of the in function to test if an element is in a slice.
func Example_inSlice() {
	out, err := af.In.ValidateRun("a", []string{"a", "b", "c"})
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%#v", out))
	// Output: true
}

// This examples shows the use of the in function checking if an IP address is in an IP range.
func Example_inIPRange() {
	addr := net.ParseIP("192.0.2.100")
	n := net.IPNet{IP: net.ParseIP("192.0.2.0"), Mask: net.CIDRMask(24, 32)}
	out, err := af.In.ValidateRun(addr, n)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%#v", out))
	// Output: true
}
