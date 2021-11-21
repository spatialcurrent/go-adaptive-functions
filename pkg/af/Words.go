// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package af

import (
	"bufio"
	"bytes"
	"reflect"
	"strings"
)

func words(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		switch in := args[0].(type) {
		case string:
			scanner := bufio.NewScanner(strings.NewReader(in))
			scanner.Split(bufio.ScanWords)
			words := make([]string, 0)
			for scanner.Scan() {
				words = append(words, scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				return words, err
			}
			return words, nil
		case []byte:
			scanner := bufio.NewScanner(bytes.NewReader(in))
			scanner.Split(bufio.ScanWords)
			words := make([]string, 0)
			for scanner.Scan() {
				words = append(words, scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				return words, err
			}
			return words, nil
		}
	}
	return nil, &ErrInvalidArguments{Function: "Words", Arguments: args}
}

var Words = Function{
	Name:    "Words",
	Aliases: []string{"words"},
	Definitions: []Definition{
		Definition{Inputs: []interface{}{uint8SliceType}, Output: reflect.Slice},
		Definition{Inputs: []interface{}{reflect.String}, Output: reflect.Slice},
	},
	Function: words,
}
