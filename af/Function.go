// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

type Function struct {
	Name        string
	Aliases     []string
	Definitions []Definition
	Function    func(args []interface{}) (interface{}, error)
}

func (f Function) IsValid(args []interface{}) bool {
	for _, d := range f.Definitions {
		if d.IsValid(args) {
			return true
		}
	}
	return false
}

func (f Function) Validate(args []interface{}) error {
	valid := f.IsValid(args)
	if !valid {
		str := "invalid arguments for " + f.Name + " with types "
		types := make([]string, 0, len(args))
		for _, a := range args {
			types = append(types, fmt.Sprint(reflect.TypeOf(a)))
		}
		str += strings.Join(types, ", ")
		return errors.New(str)
	}
	return nil
}

func (f Function) Run(args []interface{}) (interface{}, error) {
	return f.Function(args)
}
