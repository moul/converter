package converter

type ConversionFunc func(interface{}, *interface{}) error

type Conversion interface {
	// Convert func ConversionFunc
}

func Chain(left ConversionFunc, rights ...ConversionFunc) ConversionFunc {
	fn := left
	for _, right := range rights {
		fn = Pipe(fn, right)
	}
	return fn
}

func Pipe(left, right ConversionFunc) ConversionFunc {
	return func(in interface{}, output *interface{}) error {
		var tmpOutput interface{}

		if err := left(in, &tmpOutput); err != nil {
			return err
		}

		return right(tmpOutput, output)
	}
}
