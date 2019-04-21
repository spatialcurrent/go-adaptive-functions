[![CircleCI](https://circleci.com/gh/spatialcurrent/go-adaptive-functions/tree/master.svg?style=svg)](https://circleci.com/gh/spatialcurrent/go-adaptive-functions/tree/master) [![Go Report Card](https://goreportcard.com/badge/spatialcurrent/go-adaptive-functions)](https://goreportcard.com/report/spatialcurrent/go-adaptive-functions)  [![GoDoc](https://godoc.org/github.com/spatialcurrent/go-adaptive-functions?status.svg)](https://godoc.org/github.com/spatialcurrent/go-adaptive-functions) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/spatialcurrent/go-adaptive-functions/blob/master/LICENSE.md)

# go-adaptive-functions

# Description

**go-adaptive-functions** is a library of adaptive Go functions used by [go-adaptive-functions](https://github.com/spatialcurrent/go-adaptive-functions).

# Usage

**Go**

You can import **go-adaptive-functions** as a library with:

```
import (
  "github.com/spatialcurrent/go-adaptive-functions/af"
)
```

See [af](https://godoc.org/github.com/spatialcurrent/go-adaptive-functions/af) in GoDoc for information on how to use Go API.

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-adaptive-functions/blob/master/CONTRIBUTING.md) for how to get started.

# Testing

Run test using `bash scripts/test.sh`, which runs unit tests, `go vet`, `go vet with shadow`, [errcheck](https://github.com/kisielk/errcheck), [ineffassign](https://github.com/gordonklaus/ineffassign), [staticcheck](https://staticcheck.io/), and [misspell](https://github.com/client9/misspell).

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
