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

func (chain *ConverterChain) ConversionFunc() ConversionFn {
	fn := (*chain)[0].ConversionFunc
	if len(*chain) == 1 {
		return fn
	}
	for _, right := range (*chain)[1:] {
		fn = Pipe(fn, right.ConversionFunc)
	}
	return fn
}

func (chain *ConverterChain) Convert(input interface{}, output *interface{}) error {
	return chain.ConversionFunc()(input, output)
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
