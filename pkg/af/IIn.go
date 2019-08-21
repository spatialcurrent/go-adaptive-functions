// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"reflect"
	"strings"
)

func iin(args ...interface{}) (interface{}, error) {

	if len(args) != 2 {
		return false, &ErrInvalidArguments{Function: "IIn", Arguments: args}
	}

	a := args[0]
	b := args[1]

	at := reflect.TypeOf(a)
	bt := reflect.TypeOf(b)

	if at.Kind() == reflect.String {

		av := strings.ToLower(a.(string))

		if bt.Kind() == reflect.String {

			return strings.Contains(strings.ToLower(b.(string)), av), nil

		} else if bt.Kind() == reflect.Map {

			if !at.AssignableTo(bt.Key()) {
				return false, &ErrInvalidArguments{Function: "IIn", Arguments: args}
			}

			bv := reflect.ValueOf(b)

			for _, key := range bv.MapKeys() {
				if av == strings.ToLower(key.Interface().(string)) {
					return true, nil
				}
			}

			return false, nil

		} else if bt.Kind() == reflect.Struct {

			numberOfFields := bt.NumField()
			for i := 0; i < numberOfFields; i++ {
				if av == strings.ToLower(bt.Field(i).Name) {
					return true, nil
				}
			}

			numberOfMethods := bt.NumMethod()
			for i := 0; i < numberOfMethods; i++ {
				if av == strings.ToLower(bt.Method(i).Name) {
					return true, nil
				}
			}

			return false, nil

		} else if bt.Kind() == reflect.Array || bt.Kind() == reflect.Slice {

			bv := reflect.ValueOf(b)

			length := bv.Len()

			for i := 0; i < length; i++ {
				if bvs, ok := bv.Index(i).Interface().(string); ok && av == strings.ToLower(bvs) {
					return true, nil
				}
			}

			return false, nil

		}

	} else if at.AssignableTo(stringSliceType) && bt.AssignableTo(stringSliceType) {

		av := a.([]string)
		bv := b.([]string)

		if len(av) == len(bv) && len(av) == 0 {
			return false, nil
		}

		// iterate over every index in bv
		for i := range bv {
			//lint:ignore SA6005 we do not fold here as that is handled by its own function
			if strings.ToLower(bv[i]) == strings.ToLower(av[0]) && i+len(av) < len(bv) {
				match := true
				// iterate over every index in av
				for j := range av {
					//lint:ignore SA6005 we do not fold here as that is handled by its own function
					if strings.ToLower(bv[i+j]) != strings.ToLower(av[j]) {
						match = false
						break
					}
				}
				if match {
					return true, nil
				}
			}
		}

		return false, nil

	}

	return nil, &ErrInvalidArguments{Function: "IIn", Arguments: args}

}

var IIn = Function{
	Name:    "IIn",
	Aliases: []string{"iin"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{reflect.String, reflect.String}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{reflect.String, reflect.Array}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{reflect.String, reflect.Slice}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{reflect.String, reflect.Map}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{reflect.String, reflect.Struct}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{stringSliceType, stringSliceType}, Output: reflect.Bool},
	},
	Function: iin,
}
