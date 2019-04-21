// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"net"
	"reflect"
	"strings"
)

func in(args []interface{}) (interface{}, error) {

	if len(args) != 2 {
		return false, &ErrorInvalidArguments{Function: "In", Arguments: args}
	}

	a := args[0]
	b := args[1]

	at := reflect.TypeOf(a)
	bt := reflect.TypeOf(b)

	if at.Kind() == reflect.String && bt.Kind() == reflect.String {

		return strings.Contains(b.(string), b.(string)), nil

	} else if bt.AssignableTo(ipNetType) || bt.AssignableTo(ipNetPtrType) {

		if at.AssignableTo(ipType) {
			switch av := a.(type) {
			case net.IP:
				switch bv := b.(type) {
				case net.IPNet:
					return bv.Contains(av), nil
				case *net.IPNet:
					return bv.Contains(av), nil
				}
			}
		}

	} else if at.AssignableTo(stringArrayType) && bt.AssignableTo(stringArrayType) {

		av := a.([]string)
		bv := b.([]string)

		if len(av) == len(bv) && len(av) == 0 {
			return false, nil
		}

		for i := range bv {
			if bv[i] == av[0] && i+len(av) < len(bv) {
				match := true
				for j := range av {
					if bv[i+j] != av[j] {
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

	} else if at.AssignableTo(intArrayType) && bt.AssignableTo(intArrayType) {

		av := a.([]int)
		bv := b.([]int)

		if len(av) == len(bv) && len(av) == 0 {
			return false, nil
		}

		// iterate over every index in bv
		for i := range bv {
			if bv[i] == av[0] && i+len(av) < len(bv) {
				match := true
				// iterate over every index in av
				for j := range av {
					if bv[i+j] != av[j] {
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

	} else if at.AssignableTo(byteArrayType) && bt.AssignableTo(byteArrayType) {

		av := a.([]byte)
		bv := b.([]byte)

		if len(av) == len(bv) && len(av) == 0 {
			return false, nil
		}

		// iterate over every index in bv
		for i := range bv {
			if bv[i] == av[0] && i+len(av) < len(bv) {
				match := true
				// iterate over every index in av
				for j := range av {
					if bv[i+j] != av[j] {
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

	} else if bt.Kind() == reflect.Map {

		if !at.ConvertibleTo(bt.Key()) {
			return false, &ErrorInvalidArguments{Function: "In", Arguments: args}
		}

		return reflect.ValueOf(b).MapIndex(reflect.ValueOf(a).Convert(bt.Key())).IsValid(), nil

	} else if at.Kind() == reflect.String && bt.Kind() == reflect.Struct {

		av := a.(string)

		_, ok := bt.FieldByName(av)
		if ok {
			return true, nil
		}

		_, ok = bt.MethodByName(av)
		return ok, nil

	} else if bt.Kind() == reflect.Array || bt.Kind() == reflect.Slice {

		if !at.ConvertibleTo(bt.Elem()) {
			return false, &ErrorInvalidArguments{Function: "In", Arguments: args}
		}

		bv := reflect.ValueOf(b)

		length := bv.Len()

		for i := 0; i < length; i++ {
			if reflect.DeepEqual(bv.Index(i).Interface(), a) {
				return true, nil
			}
		}

		return false, nil

	}

	return nil, &ErrorInvalidArguments{Function: "In", Arguments: args}

}

var In = Function{
	Name:    "In",
	Aliases: []string{"in"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{ipType, ipNetType}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{ipType, ipNetPtrType}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{reflect.String, reflect.String}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{intArrayType, intArrayType}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{stringArrayType, stringArrayType}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{byteArrayType, byteArrayType}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{reflect.String, reflect.Array}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{reflect.String, reflect.Slice}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{reflect.String, reflect.Map}, Output: reflect.Bool},
		Definition{Inputs: []interface{}{reflect.String, reflect.Struct}, Output: reflect.Bool},
	},
	Function: in,
}
