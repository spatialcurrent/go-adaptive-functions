// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"github.com/pkg/errors"
	"reflect"
	"testing"
)

func TestIIn(t *testing.T) {

	testCases := []TestCase{
		NewTestCase([]interface{}{"A", map[string]interface{}{"a": "x"}}, IIn, true),
		NewTestCase([]interface{}{"A", []string{"a", "b"}}, IIn, true),
		NewTestCase([]interface{}{"A", []interface{}{"a"}}, IIn, true),
		NewTestCase([]interface{}{"A", []interface{}{"b", "c"}}, IIn, false),
		NewTestCase([]interface{}{"Sushi", []interface{}{"japanese", "sushi"}}, IIn, true),
		NewTestCase(
			[]interface{}{
				[]string{"My", "name"},
				[]string{"hello", "world", "my", "Name", "is"},
			},
			IIn,
			true,
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
