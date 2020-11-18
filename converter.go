package converter

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var Converters = map[string]interface{}{
	"rev":   rev,
	"upper": strings.ToUpper,
	"lower": strings.ToLower,
	"title": strings.Title,

	// internal
	"_parse-string":      parseString,
	"_bytes-to-string":   func(in []byte) string { return string(in) },
	"_string-to-bytes":   func(in string) []byte { return []byte(in) },
	"_int64-to-string":   func(in int64) string { return strconv.FormatInt(in, 10) },
	"_string-to-int64":   func(in string) (int64, error) { return strconv.ParseInt(strings.TrimSpace(in), 10, 0) },
	"_string-to-float64": func(in string) (float64, error) { return strconv.ParseFloat(strings.TrimSpace(in), 64) },
	"_float64-to-string": func(in float64) string { return strconv.FormatFloat(in, 'f', -1, 64) },
}

type Func func(interface{}) (interface{}, error)

func Chain(input interface{}, converters []string) (interface{}, error) {
	fn, err := ChainFunc(converters)
	if err != nil {
		return nil, fmt.Errorf("failed to build chain func: %w", err)
	}

	return fn(input)
}

func ChainFunc(converters []string) (Func, error) {
	if len(converters) == 0 {
		return func(input interface{}) (interface{}, error) {
			return input, nil
		}, nil
	}

	var ret Func

	for _, name := range converters {
		converter, found := Converters[name]
		if !found {
			return nil, fmt.Errorf("no such converter with name %q", name)
		}

		fn, err := converterToFunc(converter)
		if err != nil {
			return nil, fmt.Errorf(
				"invalid converter signature %q: %w",
				reflect.TypeOf(converter).String(),
				err,
			)
		}

		ret = bridgeFuncs(ret, fn)
	}

	return ret, nil
}

func bridgeFuncs(left, right Func) Func {
	if left == nil {
		return right
	}
	return func(input interface{}) (interface{}, error) {
		ret, err := left(input)
		if err != nil {
			return nil, err
		}
		return right(ret)
	}
}

func converterToFunc(converter interface{}) (Func, error) {
	x := reflect.TypeOf(converter)
	// 1 input parameter, 1 or 2 output parameters
	if x.NumIn() != 1 || x.NumOut() < 1 || x.NumOut() > 2 || x.IsVariadic() {
		return nil, fmt.Errorf("invalid amount of parameters")
	}
	var retErrT reflect.Type
	if x.NumOut() == 2 {
		retErrT = x.Out(1)
		errorInterface := reflect.TypeOf((*error)(nil)).Elem()
		if !retErrT.Implements(errorInterface) {
			return nil, fmt.Errorf("second return argument should be an error")
		}
	}

	v := reflect.ValueOf(converter)
	fn := func(input interface{}) (interface{}, error) {
		args := []reflect.Value{reflect.ValueOf(input)}
		ret := v.Call(args)

		retV := ret[0].Interface()
		var err error
		if len(ret) == 2 {
			if v := ret[1].Interface(); v != nil {
				err = v.(error)
			}
		}
		return retV, err
	}
	return fn, nil
}

// parseString takes a string in input and tries to cast it in a more specific type (date, int, etc).
// This function should be the first one to be called in a chain when using a CLI.
func parseString(in string) interface{} {
	if n, err := strconv.ParseInt(in, 10, 0); err == nil {
		return n
	}
	if n, err := strconv.ParseFloat(in, 64); err == nil {
		return n
	}

	// FIXME: try to parse other formats first
	return in
}

func rev(in string) (string, error) {
	runes := []rune(in)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes), nil
}
