// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// af.mod.js is the package for go-adaptive-functions (AF) that is built as a JavaScript module.
// In modern JavaScript, the module can be imported using destructuring assignment.
// The functions are defined in the Exports variable in the afjs package.
//
// Usage
//	// Below is a simple set of examples of how to use this package in a JavaScript application.
//
//	// load variadic functions into current scope
//	const { concat } = require('./dist/af.global.min.js);
//
//  // execute function like you would any other JavaScript function
//  // returns object of {result, err}.  err is null if no error.
//	var { result, err } = concat([["Hello", " "], "World"]);
//
// References
//	- https://godoc.org/pkg/github.com/spatialcurrent/go-adaptive-functions/pkg/afjs/
//	- https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment
package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/spatialcurrent/go-adaptive-functions/pkg/afjs"
)

func main() {
	jsModuleExports := js.Module.Get("exports")
	for name, value := range afjs.Exports {
		jsModuleExports.Set(name, value)
	}
}
