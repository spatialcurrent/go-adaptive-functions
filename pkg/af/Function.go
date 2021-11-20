// =================================================================
//
// Copyright (C) 2021 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

// Function is a struct used for giving an underlying function a name and definitions.
type Function struct {
	Name        string                                         // the name of the function
	Description string                                         // a description of the function
	Aliases     []string                                       // aliases for the function used when providing a public API
	Definitions []Definition                                   // the definitions of the function
	Function    func(args ...interface{}) (interface{}, error) // the underlying function to execute
}

// Map returns a map of metadata about the function.
func (f Function) Map() map[string]interface{} {
	return map[string]interface{}{
		"name":        f.Name,
		"description": f.Description,
		"aliases":     f.Aliases,
	}
}

// IsBoolean returns true if the function always returns a boolean value.
func (f Function) IsBoolean() bool {
	if len(f.Definitions) == 0 {
		return false
	}
	for _, d := range f.Definitions {
		if !d.IsBoolean() {
			return false
		}
	}
	return true
}

// IsInteger returns true if the function always returns an integer value.
func (f Function) IsInteger() bool {
	if len(f.Definitions) == 0 {
		return false
	}
	for _, d := range f.Definitions {
		if !d.IsInteger() {
			return false
		}
	}
	return true
}

// IsFloat returns true if the function always returns a floating point number value.
func (f Function) IsFloat() bool {
	if len(f.Definitions) == 0 {
		return false
	}
	for _, d := range f.Definitions {
		if !d.IsFloat() {
			return false
		}
	}
	return true
}

// IsNumber returns true if the function always returns a numeric value.
func (f Function) IsNumber() bool {
	if len(f.Definitions) == 0 {
		return false
	}
	for _, d := range f.Definitions {
		if !d.IsNumber() {
			return false
		}
	}
	return true
}

// IsString returns true if the function always returns a string value.
func (f Function) IsString() bool {
	if len(f.Definitions) == 0 {
		return false
	}
	for _, d := range f.Definitions {
		if !d.IsString() {
			return false
		}
	}
	return true
}

// IsValid returns true if the arguments match a definition of the function.
func (f Function) IsValid(args ...interface{}) bool {
	if len(f.Definitions) == 0 {
		return true
	}
	for _, d := range f.Definitions {
		if d.IsValid(args...) {
			return true
		}
	}
	return false
}

// Validate returns a ErrInvalidArguments error if the arguments do not match a definition of the function.
func (f Function) Validate(args ...interface{}) error {
	valid := f.IsValid(args...)
	if !valid {
		return ErrInvalidArguments{Function: f.Name, Arguments: args}
	}
	return nil
}

// Run executes the function with the provided arguments and returns the result, and error if any.
// It accepts variadic input.
func (f Function) Run(args ...interface{}) (interface{}, error) {
	return f.Function(args...)
}

// ValidateRun validates the function arguments and then runs the function if valid.
// It accepts variadic input.
func (f Function) ValidateRun(args ...interface{}) (interface{}, error) {
	err := f.Validate(args...)
	if err != nil {
		return nil, err
	}
	return f.Run(args...)
}

// MustRun executes the function with the provided arguments and returns the result.  If there is any error, then panics.
func (f Function) MustRun(args ...interface{}) interface{} {
	output, err := f.Function(args...)
	if err != nil {
		panic(err)
	}
	return output
}
