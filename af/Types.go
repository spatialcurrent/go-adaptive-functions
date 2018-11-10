// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"net"
	"reflect"
)

var interfaceType = reflect.TypeOf(map[string]interface{}{}).Elem()

var boolType = reflect.TypeOf(false)

var byteType = reflect.TypeOf(byte(0))
var stringType = reflect.TypeOf("")
var emptyStructType = reflect.TypeOf(struct{}{})

var intType = reflect.TypeOf(0)
var int8Type = reflect.TypeOf(int8(0))
var int16Type = reflect.TypeOf(int16(0))
var int32Type = reflect.TypeOf(int32(0))
var int64Type = reflect.TypeOf(int64(0))

var uint8Type = reflect.TypeOf(uint8(0))
var uint16Type = reflect.TypeOf(uint16(0))
var uint32Type = reflect.TypeOf(uint32(0))
var uint64Type = reflect.TypeOf(uint64(0))

var float32Type = reflect.TypeOf(float32(0.0))
var float64Type = reflect.TypeOf(0.0)

var byteArrayType = reflect.TypeOf([]byte{})
var stringArrayType = reflect.TypeOf([]string{})
var uint8ArrayType = reflect.TypeOf([]uint8{})
var intArrayType = reflect.TypeOf([]int{})
var int64ArrayType = reflect.TypeOf([]int64{})
var float64ArrayType = reflect.TypeOf([]float64{})
var stringIntMapSliceType = reflect.TypeOf([]map[string]int{})
var interfaceArrayType = reflect.TypeOf([]interface{}{})
var interface2DArrayType = reflect.TypeOf([][]interface{}{})

var stringSetType = reflect.TypeOf(map[string]struct{}{})
var intSetType = reflect.TypeOf(map[int]struct{}{})
var interfaceSetType = reflect.TypeOf(map[interface{}]struct{}{})

var stringIntMapType = reflect.TypeOf(map[string]int{})

var ipType = reflect.TypeOf(net.IP([]byte{}))
var ipNetType = reflect.TypeOf(net.IPNet{})
var ipNetPtrType = reflect.TypeOf(&net.IPNet{})
