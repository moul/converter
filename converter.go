package converter

import (
	"fmt"
	"reflect"
)

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
	if x.NumOut() == 2 { // nolint:gomnd
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
		if len(ret) == 2 { // nolint:gomnd
			if v := ret[1].Interface(); v != nil {
				err = v.(error)
			}
		}
		return retV, err
	}
	return fn, nil
}
