package converter

import (
	"encoding/base32"
	"encoding/base64"
)

//type Type interface{}

type Conversion interface {
	Convert(in interface{}, out *interface{}) error
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
