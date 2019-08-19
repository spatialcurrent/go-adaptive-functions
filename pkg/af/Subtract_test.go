// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestSubtract(t *testing.T) {

	testCases := []TestCase{
		NewTestCase([]interface{}{1, 2}, Subtract, -1),
		NewTestCase([]interface{}{4, 2.0}, Subtract, 2.0),
		NewTestCase(
			[]interface{}{
				map[string]string{"a": "x", "b": "y"},
				map[string]string{"a": "d"},
			},
			Subtract,
			map[string]string{"b": "y"},
		),
		NewTestCase(
			[]interface{}{
				map[string]string{"a": "x", "b": "y"},
				map[string]struct{}{"a": struct{}{}},
			},
			Subtract,
			map[string]string{"b": "y"},
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
