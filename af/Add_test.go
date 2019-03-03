// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"github.com/pkg/errors"
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {

	testCases := []TestCase{
		NewTestCase([]interface{}{1, 2}, Add, 3),
		NewTestCase([]interface{}{4, 2.0}, Add, 6.0),
		NewTestCase(
			[]interface{}{
				map[string]string{"a": "x", "b": "y"},
				map[string]string{"a": "d"},
			},
			Add,
			map[string]string{"a": "d", "b": "y"}),
		NewTestCase(
			[]interface{}{
				[]uint8{uint8(0), uint8(1), uint8(2)},
				[]uint8{uint8(3), uint8(4), uint8(5)},
			},
			Add,
			[]uint8{uint8(0), uint8(1), uint8(2), uint8(3), uint8(4), uint8(5)}),
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
