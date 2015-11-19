package converter

import "time"

func DateToTimeConverter(format string) ConversionFn {
	return func(in interface{}, out *interface{}) (err error) {
		*out, err = time.Parse(format, in.(string))
		return err
	}
}

var ConvertANSICToTime = DateToTimeConverter(time.ANSIC)
var ConvertRFC3339ToTime = DateToTimeConverter(time.RFC3339)
var ConvertRFC822ToTime = DateToTimeConverter(time.RFC822)
var ConvertRFC850ToTime = DateToTimeConverter(time.RFC850)
var ConvertRFC1123ToTime = DateToTimeConverter(time.RFC1123)
var ConvertUnixDateToTime = DateToTimeConverter(time.UnixDate)

func ConvertTimeToUnix(in interface{}, out *interface{}) (err error) {
	*out = in.(time.Time).Unix()
	return nil
}
