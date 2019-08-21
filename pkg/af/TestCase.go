// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

// TestCase is a struct containing the variables for a unit test of function evaluation
type TestCase struct {
	Inputs   []interface{} // the inputs to the function
	Function Function      // the function to evaluate
	Output   interface{}   // the result of the evaluation
}

// NewTestCase returns a new TestCase
func NewTestCase(inputs []interface{}, function Function, output interface{}) TestCase {
	return TestCase{
		Inputs:   inputs,
		Function: function,
		Output:   output,
	}
}
