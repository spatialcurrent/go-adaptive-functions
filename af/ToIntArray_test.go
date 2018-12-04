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

func TestToIntArray(t *testing.T) {

	testCases := []TestCase{
		NewTestCase([]interface{}{[]int{12, 10}}, ToIntArray, []int{12, 10}),
		NewTestCase([]interface{}{[]float64{12.0, 10.0}}, ToIntArray, []int{12, 10}),
		NewTestCase([]interface{}{[]uint8{uint8(12), uint8(10)}}, ToIntArray, []int{12, 10}),
		NewTestCase([]interface{}{[]interface{}{12, 10}}, ToIntArray, []int{12, 10}),
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