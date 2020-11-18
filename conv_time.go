package converter

import (
	"strings"
	"time"
)

// nolint:gochecknoinits // need a refactor to remove it
func init() {
	RegisterConverter(NewConverter("to-unix").SetTypes("time.Time", "int64").SetConversionFunc(ConvertTimeToUnix))
	RegisterConverter(NewConverter("parse-ansi-date").SetTypes("string", "time.Time").SetConversionFunc(ConvertANSICToTime))
	RegisterConverter(NewConverter("parse-rfc339-date").SetTypes("string", "time.Time").SetConversionFunc(ConvertRFC3339ToTime))
	RegisterConverter(NewConverter("parse-rfc822-date").SetTypes("string", "time.Time").SetConversionFunc(ConvertRFC822ToTime))
	RegisterConverter(NewConverter("parse-rfc850-date").SetTypes("string", "time.Time").SetConversionFunc(ConvertRFC850ToTime))
	RegisterConverter(NewConverter("parse-rfc1123-date").SetTypes("string", "time.Time").SetConversionFunc(ConvertRFC1123ToTime))
	RegisterConverter(NewConverter("parse-unix-date").SetTypes("string", "time.Time").SetConversionFunc(ConvertUnixDateToTime))
	RegisterConverter(NewConverter("parse-date").SetTypes("string", "time.Time").SetConversionFunc(ConvertDateToTime).SetDefaultTypeConverter())
	RegisterConverter(NewConverter("time-to-string").SetTypes("time.Time", "string").SetConversionFunc(ConvertTimeToString).SetDefaultTypeConverter())
	RegisterConverter(NewConverter("parse-unix-timestamp").SetTypes("int64", "time.Time").SetConversionFunc(ConvertUnixTimestampToTime))
}

func DateToTimeConverter(format string) ConversionFn {
	return func(in interface{}, out *interface{}) (err error) {
		*out, err = time.Parse(format, strings.TrimSpace(in.(string)))
		return err
	}
}

var (
	ConvertANSICToTime    = DateToTimeConverter(time.ANSIC)
	ConvertRFC3339ToTime  = DateToTimeConverter(time.RFC3339)
	ConvertRFC822ToTime   = DateToTimeConverter(time.RFC822)
	ConvertRFC850ToTime   = DateToTimeConverter(time.RFC850)
	ConvertRFC1123ToTime  = DateToTimeConverter(time.RFC1123)
	ConvertUnixDateToTime = DateToTimeConverter(time.UnixDate)
)

func ConvertTimeToUnix(in interface{}, out *interface{}) (err error) {
	*out = in.(time.Time).Unix()
	return nil
}

func ConvertTimeToString(in interface{}, out *interface{}) (err error) {
	*out = in.(time.Time).String()
	return err
}

func ConvertDateToTime(in interface{}, out *interface{}) (err error) {
	input := strings.TrimSpace(in.(string))
	formats := []string{
		time.ANSIC, time.RFC3339, time.RFC822, time.RFC850, time.RFC1123, time.UnixDate,
	}
	for _, format := range formats {
		if *out, err = time.Parse(format, input); err == nil {
			return nil
		}
	}
	*out = nil
	return err
}

func ConvertUnixTimestampToTime(in interface{}, out *interface{}) (err error) {
	*out = time.Unix(in.(int64), 0)
	return nil
}
