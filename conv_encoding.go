package converter

import (
	"bytes"
	"encoding/base32"
	"encoding/base64"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/mr-tron/base58"
)

// nolint:gochecknoinits // need a refactor to remove it
func init() {
	RegisterConverter(NewConverter("base58-encode").SetTypes("[]byte", "string").SetConversionFunc(ConvertBytesToBase58))
	RegisterConverter(NewConverter("base58-decode").SetTypes("string", "[]byte").SetConversionFunc(ConvertBase58ToBytes))
	RegisterConverter(NewConverter("base64-encode").SetTypes("[]byte", "string").SetConversionFunc(ConvertBytesToBase64))
	RegisterConverter(NewConverter("base64-decode").SetTypes("string", "[]byte").SetConversionFunc(ConvertBase64ToBytes))
	RegisterConverter(NewConverter("urlbase64-encode").SetTypes("[]byte", "string").SetConversionFunc(ConvertBytesToURLBase64))
	RegisterConverter(NewConverter("urlbase64-decode").SetTypes("string", "[]byte").SetConversionFunc(ConvertURLBase64ToBytes))
	RegisterConverter(NewConverter("rawurlbase64-encode").SetTypes("[]byte", "string").SetConversionFunc(ConvertBytesToRawURLBase64))
	RegisterConverter(NewConverter("rawurlbase64-decode").SetTypes("string", "[]byte").SetConversionFunc(ConvertRawURLBase64ToBytes))
	RegisterConverter(NewConverter("base32-encode").SetTypes("[]byte", "string").SetConversionFunc(ConvertBytesToBase32))
	RegisterConverter(NewConverter("base32-decode").SetTypes("string", "[]byte").SetConversionFunc(ConvertBase32ToBytes))
	RegisterConverter(NewConverter("hex-encode").SetTypes("[]byte", "string").SetConversionFunc(ConvertBytesToHex))
	RegisterConverter(NewConverter("hex-decode").SetTypes("string", "[]byte").SetConversionFunc(ConvertHexToBytes))
	RegisterConverter(NewConverter("xml-encode").SetTypes("interface{}", "[]byte").SetConversionFunc(ConvertStructToXML))
	RegisterConverter(NewConverter("xml-decode").SetTypes("[]byte", "interface{}").SetConversionFunc(ConvertXMLToStruct))
	RegisterConverter(NewConverter("json-encode").SetTypes("interface{}", "[]byte").SetConversionFunc(ConvertStructToJSON))
	RegisterConverter(NewConverter("json-decode").SetTypes("[]byte", "interface{}").SetConversionFunc(ConvertJSONToStruct))
	RegisterConverter(NewConverter("toml-encode").SetTypes("[]byte", "interface{}").SetConversionFunc(ConvertStructToTOML))
	RegisterConverter(NewConverter("csv-decode").SetTypes("string", "[][]string").SetConversionFunc(ConvertStringToCsv))
}

func ConvertBytesToRawURLBase64(in interface{}, out *interface{}) error {
	*out = base64.RawURLEncoding.EncodeToString(in.([]byte))
	return nil
}

func ConvertRawURLBase64ToBytes(in interface{}, out *interface{}) (err error) {
	*out, err = base64.RawURLEncoding.DecodeString(in.(string))
	return err
}

func ConvertBytesToURLBase64(in interface{}, out *interface{}) error {
	*out = base64.URLEncoding.EncodeToString(in.([]byte))
	return nil
}

func ConvertURLBase64ToBytes(in interface{}, out *interface{}) (err error) {
	*out, err = base64.URLEncoding.DecodeString(in.(string))
	return err
}

func ConvertBytesToBase64(in interface{}, out *interface{}) error {
	*out = base64.StdEncoding.EncodeToString(in.([]byte))
	return nil
}

func ConvertBase64ToBytes(in interface{}, out *interface{}) (err error) {
	*out, err = base64.StdEncoding.DecodeString(in.(string))
	return err
}

func ConvertBytesToBase58(in interface{}, out *interface{}) error {
	*out = base58.Encode(in.([]byte))
	return nil
}

func ConvertBase58ToBytes(in interface{}, out *interface{}) (err error) {
	*out, err = base58.Decode(in.(string))
	return err
}

func ConvertBytesToBase32(in interface{}, out *interface{}) error {
	*out = base32.StdEncoding.EncodeToString(in.([]byte))
	return nil
}

func ConvertBase32ToBytes(in interface{}, out *interface{}) (err error) {
	*out, err = base32.StdEncoding.DecodeString(in.(string))
	return err
}

func ConvertBytesToHex(in interface{}, out *interface{}) error {
	*out = hex.EncodeToString(in.([]byte))
	return nil
}

func ConvertHexToBytes(in interface{}, out *interface{}) (err error) {
	*out, err = hex.DecodeString(in.(string))
	return err
}

func ConvertXMLToStruct(in interface{}, out *interface{}) (err error) {
	return xml.Unmarshal(in.([]byte), out)
}

func ConvertStructToXML(in interface{}, out *interface{}) (err error) {
	*out, err = xml.Marshal(in)
	return err
}

func ConvertJSONToStruct(in interface{}, out *interface{}) (err error) {
	return json.Unmarshal(in.([]byte), out)
}

func ConvertStructToJSON(in interface{}, out *interface{}) (err error) {
	*out, err = json.Marshal(in)
	return err
}

func ConvertStructToTOML(in interface{}, out *interface{}) (err error) {
	buf := new(bytes.Buffer)
	err = toml.NewEncoder(buf).Encode(in)
	*out = buf.Bytes()
	return err
}

func ConvertStringToCsv(in interface{}, out *interface{}) (err error) {
	r := csv.NewReader(strings.NewReader(in.(string)))
	*out, err = r.ReadAll()
	return err
}
