package converter

import (
	"strings"
	"time"
)

func init() {
	RegisterConverter(NewConverter("to-unix").SetTypes("time.Time", "int64").SetConversionFunc(ConvertTimeToUnix))
	RegisterConverter(NewConverter("parse-ansi-date").SetTypes("string", "time.Time").SetConversionFunc(ConvertANSICToTime))
	RegisterConverter(NewConverter("parse-rfc339-date").SetTypes("string", "time.Time").SetConversionFunc(ConvertRFC3339ToTime))
	RegisterConverter(NewConverter("parse-rfc822-date").SetTypes("string", "time.Time").SetConversionFunc(ConvertRFC822ToTime))
	RegisterConverter(NewConverter("parse-rfc850-date").SetTypes("string", "time.Time").SetConversionFunc(ConvertRFC850ToTime))
	RegisterConverter(NewConverter("parse-rfc1123-date").SetTypes("string", "time.Time").SetConversionFunc(ConvertRFC1123ToTime))
	RegisterConverter(NewConverter("parse-unix-date").SetTypes("string", "time.Time").SetConversionFunc(ConvertUnixDateToTime))
}

func DateToTimeConverter(format string) ConversionFn {
	return func(in interface{}, out *interface{}) (err error) {
		*out, err = time.Parse(format, strings.TrimSpace(in.(string)))
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
