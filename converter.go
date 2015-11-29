package converter

import "fmt"

type ConversionFn func(interface{}, *interface{}) error
type StreamConvFn func(chan interface{}) chan interface{}

type ConverterChain []*Converter

func NewConverterChain(filterNames []string) (ConverterChain, error) {
	if len(filterNames) == 0 {
		return nil, fmt.Errorf("you should have at least one filter")
	}
	chain := ConverterChain{}
	for _, name := range filterNames {
		converter, err := GetConverterByName(name)
		if err != nil {
			return nil, err
		}
		chain = append(chain, converter)
	}
	return chain, nil
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

func (chain *ConverterChain) ConversionFunc(inType, outType string) (ConversionFn, error) {
	if len(*chain) == 0 {
		return nil, fmt.Errorf("you should have at least one converter")
	}

	fn := (*chain)[0].ConversionFunc
	if convertFn := GetTypeConversionFunc(inType, (*chain)[0].InputType); convertFn != nil {
		fn = Pipe(convertFn, fn)
	}

	if len(*chain) == 1 {
		return fn, nil
	}

	inType = (*chain)[0].OutputType
	for _, right := range (*chain)[1:] {
		if convertFn := GetTypeConversionFunc(inType, right.InputType); convertFn != nil {
			fn = Pipe(fn, convertFn)
		}
		fn = Pipe(fn, right.ConversionFunc)
		inType = right.OutputType
	}
	return fn, nil
}

func (chain *ConverterChain) Convert(input interface{}, output *interface{}) error {
	fn, err := chain.ConversionFunc("interface{}", "interface{}")
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
	return nil, fmt.Errorf("no such filter %q", name)
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
			for {
				select {
				case input := <-in:
					var output interface{}
					_ = conversionFn(input, &output)
					// FIXME: check err
					out <- output
				}
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
