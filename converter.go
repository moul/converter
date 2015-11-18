package converter

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"strconv"
)

type ConversionFunc func(interface{}, *interface{}) error

type Conversion interface {
	// Convert func ConversionFunc
}

func Pipe(left, right ConversionFunc) ConversionFunc {
	return func(in interface{}, output *interface{}) error {
		var tmpOutput interface{}

		err := left(in, &tmpOutput)
		if err != nil {
			return err
		}

		return right(tmpOutput, output)
	}
}

func ConvertBytesToBase64(in interface{}, out *interface{}) error {
	*out = base64.StdEncoding.EncodeToString(in.([]byte))
	return nil
}

func ConvertBase64ToBytes(in interface{}, out *interface{}) error {
	data, err := base64.StdEncoding.DecodeString(in.(string))
	if err != nil {
		return err
	}
	*out = data
	return nil
}

func ConvertBytesToBase32(in interface{}, out *interface{}) error {
	*out = base32.StdEncoding.EncodeToString(in.([]byte))
	return nil
}

func ConvertBase32ToBytes(in interface{}, out *interface{}) error {
	data, err := base32.StdEncoding.DecodeString(in.(string))
	if err != nil {
		return err
	}
	*out = data
	return nil
}

func ConvertBytesToString(in interface{}, out *interface{}) error {
	*out = string(in.([]byte))
	return nil
}

func ConvertStringToBytes(in interface{}, out *interface{}) error {
	*out = []byte(in.(string))
	return nil
}

func ConvertBytesToHex(in interface{}, out *interface{}) error {
	*out = hex.EncodeToString(in.([]byte))
	return nil
}

func ConvertHexToBytes(in interface{}, out *interface{}) error {
	data, err := hex.DecodeString(in.(string))
	if err != nil {
		return err
	}
	*out = data
	return nil
}

func ConvertIntegerToString(in interface{}, out *interface{}) error {
	*out = strconv.Itoa(in.(int))
	return nil
}

func ConvertStringToInteger(in interface{}, out *interface{}) (err error) {
	*out, err = strconv.Atoi(in.(string))
	return err
}

func ConvertStringToFloat(in interface{}, out *interface{}) (err error) {
	*out, err = strconv.ParseFloat(in.(string), 64)
	return err
}

func ConvertFloatToString(in interface{}, out *interface{}) (err error) {
	*out = strconv.FormatFloat(in.(float64), 'f', -1, 64)
	return err
}
