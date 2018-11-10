// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"github.com/pkg/errors"
	"net"
	"reflect"
	"testing"
)

func TestIn(t *testing.T) {

	testCases := []TestCase{
		NewTestCase([]interface{}{"a", map[string]interface{}{"a": "x"}}, In, true),
		NewTestCase([]interface{}{"a", []string{"a"}}, In, true),
		NewTestCase([]interface{}{"a", []interface{}{"a"}}, In, true),
		NewTestCase([]interface{}{"a", []interface{}{"b", "c"}}, In, false),
		NewTestCase(
			[]interface{}{
				[]string{"my", "name"},
				[]string{"hello", "world", "my", "name", "is"},
			},
			In,
			true,
		),
		NewTestCase(
			[]interface{}{
				[]int{9, 8},
				[]int{10, 9, 8, 7},
			},
			In,
			true,
		),
		NewTestCase(
			[]interface{}{
				net.ParseIP("192.0.2.100"),
				net.IPNet{IP: net.ParseIP("192.0.2.0"), Mask: net.CIDRMask(24, 32)},
			},
			In,
			true,
		),
		NewTestCase(
			[]interface{}{
				net.ParseIP("192.0.2.100"),
				&net.IPNet{IP: net.ParseIP("192.0.2.0"), Mask: net.CIDRMask(24, 32)},
			},
			In,
			true,
		),
		NewTestCase(
			[]interface{}{
				net.ParseIP("192.0.2.100"),
				net.IPNet{IP: net.ParseIP("192.0.2.0"), Mask: net.CIDRMask(31, 32)},
			},
			In,
			false,
		),
	}

	for _, testCase := range testCases {

		valid := testCase.Function.IsValid(testCase.Inputs)
		if !valid {
			t.Errorf("inputs (%v) to function %q are invalid", testCase.Inputs, testCase.Function.Name)
		}
		got, err := testCase.Function.Run(testCase.Inputs)
		if err != nil {
			t.Errorf(errors.Wrap(err, "error running function \""+reflect.TypeOf(testCase.Function).Name()+"\"").Error())
		} else if !reflect.DeepEqual(got, testCase.Output) {
			t.Errorf(testCase.Function.Name+"(%v) == %v (%v), want %v (%s)", testCase.Inputs, got, reflect.TypeOf(got), testCase.Output, reflect.TypeOf(testCase.Output))
		}
	}

}
