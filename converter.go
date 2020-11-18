package converter

import (
	"errors"
	"fmt"
)

type (
	ConversionFn func(interface{}, *interface{}) error
	StreamConvFn func(chan interface{}) chan interface{}
	Flow         []*Converter
)

var ErrMinOneFilter = errors.New("you should have at least one filter")

func NewFlow(filterNames []string) (Flow, error) {
	if len(filterNames) == 0 {
		return nil, ErrMinOneFilter
	}
	flow := Flow{}
	for _, name := range filterNames {
		converter, err := GetConverterByName(name)
		if err != nil {
			return nil, err
		}
		flow = append(flow, converter)
	}
	return flow, nil
}

func GetTypeConversionFunc(inType, outType string) ConversionFn {
	if inType == outType {
		return nil
	}
	if inType == "interface{}" || outType == "interface{}" {
		return nil
	}
	for _, converter := range RegisteredConverters {
		if converter.InputType == inType && converter.OutputType == outType && converter.IsDefaultTypeConverter {
			return converter.ConversionFunc
		}
	}
	return nil
}

func (flow *Flow) ConversionFunc(inType, outType string) (ConversionFn, error) {
	if len(*flow) == 0 {
		return nil, ErrMinOneFilter
	}

	lastRealInType := inType

	fn := (*flow)[0].ConversionFunc
	if convertFn := GetTypeConversionFunc(inType, (*flow)[0].InputType); convertFn != nil {
		fn = Pipe(convertFn, fn)
	}

	if len(*flow) == 1 {
		return fn, nil
	}

	inType = (*flow)[0].OutputType
	for _, right := range (*flow)[1:] {
		if inType != "interface{}" {
			lastRealInType = inType
		}
		if convertFn := GetTypeConversionFunc(lastRealInType, right.InputType); convertFn != nil {
			fn = Pipe(fn, convertFn)
		}
		fn = Pipe(fn, right.ConversionFunc)
		inType = right.OutputType
	}
	return fn, nil
}

func (flow *Flow) Convert(input interface{}, output *interface{}) error {
	fn, err := flow.ConversionFunc("interface{}", "interface{}")
	if err != nil {
		return err
	}
	return fn(input, output)
}

func GetConverterByName(name string) (*Converter, error) {
	for _, converter := range RegisteredConverters {
		if converter.Name == name {
			return converter, nil
		}
	}
	return nil, fmt.Errorf("no such filter %q", name) // nolint:goerr113
}

func Chain(left ConversionFn, rights ...ConversionFn) ConversionFn {
	fn := left
	for _, right := range rights {
		fn = Pipe(fn, right)
	}
	return fn
}

func Pipe(left, right ConversionFn) ConversionFn {
	return func(in interface{}, output *interface{}) error {
		var tmpOutput interface{}
		if err := left(in, &tmpOutput); err != nil {
			return err
		}
		return right(tmpOutput, output)
	}
}

func ConversionToStreamConv(conversionFn ConversionFn) StreamConvFn {
	return func(in chan interface{}) chan interface{} {
		out := make(chan interface{})
		go func() {
			for input := range in {
				var output interface{}
				_ = conversionFn(input, &output)
				// FIXME: check err
				out <- output
			}
		}()
		return out
	}
}

func StreamPipe(left, right StreamConvFn) StreamConvFn {
	return func(in chan interface{}) chan interface{} {
		return right(left(in))
	}
}

func StreamChain(streamFuncs ...StreamConvFn) StreamConvFn {
	return func(in chan interface{}) chan interface{} {
		left := in
		for _, right := range streamFuncs {
			left = right(left)
		}
		return left
	}
}

type Converter struct {
	InputType              string
	OutputType             string
	Name                   string
	ConversionFunc         ConversionFn
	StreamConvFunc         StreamConvFn
	IsDefaultTypeConverter bool
}

func (conv *Converter) SetType(ioType string) *Converter {
	conv.InputType = ioType
	conv.OutputType = ioType
	return conv
}

func (conv *Converter) SetTypes(inType, outType string) *Converter {
	conv.InputType = inType
	conv.OutputType = outType
	return conv
}

func (conv *Converter) SetConversionFunc(fn ConversionFn) *Converter {
	conv.ConversionFunc = fn
	return conv
}

func (conv *Converter) SetStreamConvFunc(fn StreamConvFn) *Converter {
	conv.StreamConvFunc = fn
	return conv
}

func (conv *Converter) SetDefaultTypeConverter() *Converter {
	conv.IsDefaultTypeConverter = true
	return conv
}

func NewConverter(name string) *Converter {
	return &Converter{
		InputType:  "interface{}",
		OutputType: "interface{}",
		Name:       name,
	}
}

var RegisteredConverters []*Converter

func RegisterConverter(converter *Converter) {
	RegisteredConverters = append(RegisteredConverters, converter)
}
