// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// Package afjs includes functions for the JavaScript build of go-adaptive-functions.
//
package afjs

import (
	"github.com/gopherjs/gopherjs/js"

	"github.com/spatialcurrent/go-adaptive-functions/pkg/af"
)

func isArray(object *js.Object) bool {
	return js.Global.Get("Array").Call("isArray", object).Bool()
}

func toArray(object *js.Object) []interface{} {
	arr := make([]interface{}, 0, object.Length())
	for i := 0; i < object.Length(); i++ {
		arr = append(arr, parseObject(object.Index(i)))
	}
	return arr
}

func parseObject(object *js.Object) interface{} {
	if isArray(object) {
		return toArray(object)
	}
	return object.Interface()
}

func exports() map[string]interface{} {
	m := map[string]interface{}{}
	for _, fn := range af.Functions {
		func(fn af.Function) {
			for _, alias := range fn.Aliases {
				m[alias] = func(objects ...*js.Object) map[string]interface{} {
					args := make([]interface{}, 0, len(objects))
					for _, obj := range objects {
						args = append(args, parseObject(obj))
					}
					result, err := fn.ValidateRun(args...)
					if err != nil {
						return map[string]interface{}{"result": nil, "err": err.Error()}
					}
					return map[string]interface{}{"result": result, "err": nil}
				}
			}
		}(fn)
	}
	return m
}

var Exports = exports()
