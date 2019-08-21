// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

const { sum } = require('./../../dist/af.mod.min.js');

console.log("************************************");
console.log();
// define input
const input = [1, [2, 3], 4];
console.log('Input:');
console.log(input);
console.log();
// execute like any other JavaScript function and destructure return value
var { result, err } = sum(input);
console.log('Result:');
console.log(result);
console.log();
console.log('Error:');
console.log(err);
console.log();
console.log("************************************");
console.log();
