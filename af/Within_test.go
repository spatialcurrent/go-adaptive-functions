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

func TestWithin(t *testing.T) {

	testCases := []TestCase{
		NewTestCase([]interface{}{[]float64{-80.0, 36.0}, []float64{-168.39843, 13.58192, -52.73437, 72.1818}}, Within, true),
		NewTestCase([]interface{}{[]float64{-180.0, 36.0}, []float64{-168.39843, 13.58192, -52.73437, 72.1818}}, Within, false),
		NewTestCase([]interface{}{[]float64{-80.0, 85.0}, []float64{-168.39843, 13.58192, -52.73437, 72.1818}}, Within, false),
		NewTestCase([]interface{}{[]float64{-80.0, 36, 10.0}, []float64{-168.39843, 13.58192, 0.0, -52.73437, 72.1818, 100.0}}, Within, true),
		NewTestCase([]interface{}{[]float64{-80.0, 36, 10.0}, []float64{-168.39843, 13.58192, 0.0, -52.73437, 72.1818, 8.0}}, Within, false),
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