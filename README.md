[![CircleCI](https://circleci.com/gh/spatialcurrent/go-adaptive-functions/tree/master.svg?style=svg)](https://circleci.com/gh/spatialcurrent/go-adaptive-functions/tree/master) [![Go Report Card](https://goreportcard.com/badge/spatialcurrent/go-adaptive-functions)](https://goreportcard.com/report/spatialcurrent/go-adaptive-functions)  [![GoDoc](https://godoc.org/github.com/spatialcurrent/go-adaptive-functions?status.svg)](https://godoc.org/github.com/spatialcurrent/go-adaptive-functions) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/spatialcurrent/go-adaptive-functions/blob/master/LICENSE.md)

# go-adaptive-functions

# Description

**go-adaptive-functions** is a library of adaptive Go functions used by [gotmpl](https://github.com/spatialcurrent/gotmpl), [go-dfl](https://github.com/spatialcurrent/go-dfl) and [railgun](https://github.com/spatialcurrent/railgun).

# Usage

**Go**

You can import **go-adaptive-functions** as a library with:

```
import (
  "github.com/spatialcurrent/go-adaptive-functions/pkg/af"
)
```

See [af](https://godoc.org/github.com/spatialcurrent/go-adaptive-functions/pkg/af) in GoDoc for API documentation and examples.

**Node**

AF is built as a module.  In modern JavaScript, the module can be imported using [destructuring assignment](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment).

```javascript
const { sum, min, max } = require('./dist/af.mod.min.js');
```

In legacy JavaScript, you can use the `af.global.js` file that simply adds `af` to the global scope.

# Releases

**go-adaptive-functions** is currently in **alpha**.  See releases at https://github.com/spatialcurrent/go-adaptive-functions/releases.  See the **Building** section below to build from scratch.

**JavaScript**

- `af.global.js`, `af.global.js.map` - JavaScript global build with source map
- `af.global.min.js`, `af.global.min.js.map` - Minified JavaScript global build with source map
- `af.mod.js`, `af.mod.js.map` - JavaScript module build  with source map
- `af.mod.min.js`, `af.mod.min.js.map` - Minified JavaScript module with source map

# Examples

**Go**

See the examples in [GoDoc](https://godoc.org/github.com/spatialcurrent/go-adaptive-functions/pkg/af).

**JavaScript**

See the `examples/js/index.js` file.  You can run the example with `make run_example_javascript`.

# Building

Use `make help` to see help information for each target.

**JavaScript**

You can compile GSS to pure JavaScript with the `make build_javascript` script.

# Testing

**Go**

To run Go tests use `make test_go` (or `bash scripts/test.sh`), which runs unit tests, `go vet`, `go vet with shadow`, [errcheck](https://github.com/kisielk/errcheck), [ineffassign](https://github.com/gordonklaus/ineffassign), [staticcheck](https://staticcheck.io/), and [misspell](https://github.com/client9/misspell).

**JavaScript**

To run JavaScript tests, first install [Jest](https://jestjs.io/) using `make deps_javascript`, use [Yarn](https://yarnpkg.com/en/), or another method.  Then, build the JavaScript module with `make build_javascript`.  To run tests, use `make test_javascript`.  You can also use the scripts in the `package.json`.

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-adaptive-functions/blob/master/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
