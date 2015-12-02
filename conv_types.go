package converter

import (
	"strconv"
	"strings"
)

func init() {
	RegisterConverter(NewConverter("bytes-to-string").SetTypes("[]byte", "string").SetConversionFunc(ConvertBytesToString).SetDefaultTypeConverter())
	RegisterConverter(NewConverter("string-to-bytes").SetTypes("string", "[]byte").SetConversionFunc(ConvertStringToBytes).SetDefaultTypeConverter())
	RegisterConverter(NewConverter("int-to-string").SetTypes("int64", "string").SetConversionFunc(ConvertIntegerToString).SetDefaultTypeConverter())
	RegisterConverter(NewConverter("string-to-int").SetTypes("string", "int64").SetConversionFunc(ConvertStringToInteger).SetDefaultTypeConverter())
	RegisterConverter(NewConverter("string-to-float").SetTypes("string", "float64").SetConversionFunc(ConvertStringToFloat).SetDefaultTypeConverter())
	RegisterConverter(NewConverter("float-to-string").SetTypes("float64", "string").SetConversionFunc(ConvertFloatToString).SetDefaultTypeConverter())
}

func ConvertBytesToString(in interface{}, out *interface{}) error {
	*out = string(in.([]byte))
	return nil
}

func ConvertStringToBytes(in interface{}, out *interface{}) error {
	*out = []byte(in.(string))
	return nil
}

func ConvertIntegerToString(in interface{}, out *interface{}) error {
	*out = strconv.FormatInt(in.(int64), 10)
	return nil
}

func ConvertStringToInteger(in interface{}, out *interface{}) (err error) {
	*out, err = strconv.ParseInt(strings.TrimSpace(in.(string)), 10, 0)
	return err
}

func ConvertStringToFloat(in interface{}, out *interface{}) (err error) {
	*out, err = strconv.ParseFloat(strings.TrimSpace(in.(string)), 64)
	return err
}

func ConvertFloatToString(in interface{}, out *interface{}) (err error) {
	*out = strconv.FormatFloat(in.(float64), 'f', -1, 64)
	return err
}
