package af

import (
	"fmt"
	"reflect"
	"strings"
)

// ErrorInvalidArguments is an error returned when a function is called with invalid arguments.
type ErrorInvalidArguments struct {
	Function  string        // the name of the Function
	Arguments []interface{} // the arguments for the function
}

// Error returns the error as a string.
func (e ErrorInvalidArguments) Error() string {
	str := "invalid arguments for function " + e.Function + " with types "
	types := make([]string, 0, len(e.Arguments))
	for _, a := range e.Arguments {
		types = append(types, fmt.Sprint(reflect.TypeOf(a)))
	}
	str += strings.Join(types, ", ")
	return str
}
