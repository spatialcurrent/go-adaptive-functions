// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

// Function is a struct used for giving an underlying function a name and definitions.
type Function struct {
	Name        string                                        // the name of the function
	Aliases     []string                                      // aliases for the function used when providing a public API
	Definitions []Definition                                  // the definitions of the function
	Function    func(args []interface{}) (interface{}, error) // the underlying function to execute
}

// IsValid returns true if the arguments match a definition of the function.
func (f Function) IsValid(args []interface{}) bool {
	if len(f.Definitions) == 0 {
		return true
	}
	for _, d := range f.Definitions {
		if d.IsValid(args) {
			return true
		}
	}
	return false
}

// Validate returns a ErrorInvalidArguments error if the arguments do not match a definition of the function.
func (f Function) Validate(args []interface{}) error {
	valid := f.IsValid(args)
	if !valid {
		return ErrorInvalidArguments{Function: f.Name, Arguments: args}
	}
	return nil
}

// Run executes the function with the provided arguments and returns the result, and error if any.
func (f Function) Run(args []interface{}) (interface{}, error) {
	return f.Function(args)
}

// ValidateRun validates the function arguments and then runs the function if valid.
func (f Function) ValidateRun(args []interface{}) (interface{}, error) {
	err := f.Validate(args)
	if err != nil {
		return nil, err
	}
	return f.Run(args)
}

// MustRun executes the function with the provided arguments and returns the result.  If there is any error, then panics.
func (f Function) MustRun(args []interface{}) interface{} {
	output, err := f.Function(args)
	if err != nil {
		panic(err)
	}
	return output
}
