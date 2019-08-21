# =================================================================
#
# Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
# Released as open source under the MIT License.  See LICENSE file.
#
# =================================================================

.PHONY: help
help:  ## Print the help documentation
	@grep -E '^[a-zA-Z_-\]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

#
# Dependencies
#

.PHONY: deps_go
deps_go:  ## Install Go dependencies
	go get -d -t ./...

.PHONY: deps_go_test
deps_go_test: ## Download Go dependencies for tests
	go get golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow # download shadow
	go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow # install shadow
	go get -u github.com/kisielk/errcheck # download and install errcheck
	go get -u github.com/client9/misspell/cmd/misspell # download and install misspell
	go get -u github.com/gordonklaus/ineffassign # download and install ineffassign
	go get -u honnef.co/go/tools/cmd/staticcheck # download and instal staticcheck
	go get -u golang.org/x/tools/cmd/goimports # download and install goimports

deps_gopherjs:  ## Install GopherJS
	go get -u github.com/gopherjs/gopherjs

deps_javascript:  ## Install dependencies for JavaScript tests
	npm install .

#
# Go building, formatting, testing, and installing
#


.PHONY: fmt
fmt:  ## Format Go source code
	go fmt $$(go list ./... )

.PHONY: imports
imports: ## Update imports in Go source code
	goimports -w -local github.com/spatialcurrent/go-adaptive-functions $$(find . -iname '*.go')

.PHONY: vet
vet: ## Vet Go source code
	go vet $$(go list ./... )

.PHONY: test_go
test_go: ## Run Go tests
	bash scripts/test.sh

build: build_javascript  ## Build JavaScript

#
# JavaScript
#

dist/af.mod.js:  ## Build JavaScript module
	gopherjs build -o dist/af.mod.js github.com/spatialcurrent/go-adaptive-functions/cmd/af.mod.js

dist/af.mod.min.js:  ## Build minified JavaScript module
	gopherjs build -m -o dist/af.mod.min.js github.com/spatialcurrent/go-adaptive-functions/cmd/af.mod.js

dist/af.global.js:  ## Build JavaScript library that attaches to global or window.
	gopherjs build -o dist/af.global.js github.com/spatialcurrent/go-adaptive-functions/cmd/af.global.js

dist/af.global.min.js:  ## Build minified JavaScript library that attaches to global or window.
	gopherjs build -m -o dist/af.global.min.js github.com/spatialcurrent/go-adaptive-functions/cmd/af.global.js

build_javascript: dist/af.mod.js dist/af.mod.min.js dist/af.global.js dist/af.global.min.js  ## Build artifacts for JavaScript

test_javascript:  ## Run JavaScript tests
	npm run test

lint:  ## Lint JavaScript source code
	npm run lint

#
# Examples
#

run_example_javascript: dist/af.mod.js  ## Run JavaScript module example
	node examples/js/index.mod.js

## Clean

clean:  ## Clean artifacts
	rm -fr bin
	rm -fr dist
