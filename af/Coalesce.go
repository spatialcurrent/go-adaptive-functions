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

func coalesce(args []interface{}) (interface{}, error) {

	if len(args) == 0 {
		return nil, &ErrorInvalidArguments{Function: "Coalesce", Arguments: args}
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
		v := values.Index(i)
		if v.IsValid() && (!v.IsNil()) {
			vi := v.Interface()
			vk := reflect.TypeOf(vi).Kind()
			if vk == reflect.Array || vk == reflect.Slice || vk == reflect.Map {
				if viv := reflect.ValueOf(vi); viv.IsNil() || viv.Len() == 0 {
					continue
				}
			} else if vk == reflect.String {
				if viv := reflect.ValueOf(vi); viv.Len() == 0 {
					continue
				}
			}
			return vi, nil
		}
	}

	return nil, nil
}

var Coalesce = Function{
	Name:        "Coalesce",
	Aliases:     []string{"coalesce"},
	Definitions: []Definition{},
	Function:    coalesce,
}
