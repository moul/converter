package converter

type ConversionFn func(interface{}, *interface{}) error
type StreamConvFn func(chan interface{}) chan interface{}

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
