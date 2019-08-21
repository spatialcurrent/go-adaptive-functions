// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// af.global.js is the package for go-adaptive-functions (AF) that adds AF functions to the global scope under the "af" variable.
//
// In Node, depending on where require is called and the build system used, the functions may need to be required at the top of each module file.
// In a web browser, AF can be made available to the entire web page.
// The functions are defined in the Exports variable in the afjs package.
//
// Usage
//	// Below is a simple set of examples of how to use this package in a JavaScript application.
//
//	// load functions into global scope
//	// require('./dist/af.global.min.js);
//
//  // execute function like you would any other JavaScript function
//  // returns object of {result, err}.  err is null if no error.
//	var { result, err } = af.concat([["Hello", " "], "World"]);
//
// References
//	- https://godoc.org/pkg/github.com/spatialcurrent/go-adaptive-functions/pkg/afjs/
//	- https://nodejs.org/api/globals.html#globals_global_objects
//	- https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects
package main

import (
	"github.com/gopherjs/gopherjs/js"

	"github.com/spatialcurrent/go-adaptive-functions/pkg/afjs"
)

func main() {
	js.Global.Set("af", afjs.Exports)
}
