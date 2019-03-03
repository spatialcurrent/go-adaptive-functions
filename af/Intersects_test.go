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

func TestIntersects(t *testing.T) {

	a := map[string]struct{}{"a": struct{}{}, "b": struct{}{}, "c": struct{}{}}
	b := map[string]struct{}{"a": struct{}{}, "b": struct{}{}}
	c := map[string]struct{}{"d": struct{}{}, "e": struct{}{}}

	testCases := []TestCase{
		NewTestCase([]interface{}{a, b}, Intersects, true),
		NewTestCase([]interface{}{a, c}, Intersects, false),
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
