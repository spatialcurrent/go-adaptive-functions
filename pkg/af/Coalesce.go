// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"reflect"
)

func coalesce(args ...interface{}) (interface{}, error) {

	if len(args) == 0 {
		return nil, &ErrInvalidArguments{Function: "Coalesce", Arguments: args}
	}

	var values reflect.Value
	if len(args) == 1 {
		arr := args[0]
		if t := reflect.TypeOf(arr); t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
			values = reflect.ValueOf(arr)
		} else {
			return args[0], nil
		}
	} else {
		values = reflect.ValueOf(args)
	}
	l := values.Len()
	for i := 0; i < l; i++ {
		viv := reflect.ValueOf(values.Index(i).Interface())
		if !viv.IsValid() {
			continue
		}
		vk := viv.Type().Kind()
		if vk == reflect.String {
			if viv.Len() == 0 {
				continue
			}
			return viv.Interface(), nil
		}
		if vk == reflect.Bool || vk == reflect.Int || vk == reflect.Float64 {
			return viv.Interface(), nil
		}
		if viv.IsNil() {
			continue
		}
		if vk == reflect.Array || vk == reflect.Slice || vk == reflect.Map {
			if viv.Len() == 0 {
				continue
			}
		}
		return viv.Interface(), nil
	}
	return nil, nil
}

var Coalesce = Function{
	Name:        "Coalesce",
	Aliases:     []string{"coalesce"},
	Definitions: []Definition{},
	Function:    coalesce,
}
