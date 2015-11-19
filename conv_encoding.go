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
)

func ConvertBytesToBase64(in interface{}, out *interface{}) error {
	*out = base64.StdEncoding.EncodeToString(in.([]byte))
	return nil
}

func ConvertBase64ToBytes(in interface{}, out *interface{}) (err error) {
	*out, err = base64.StdEncoding.DecodeString(in.(string))
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

func ConvertXmlToStruct(in interface{}, out *interface{}) (err error) {
	return xml.Unmarshal(in.([]byte), out)
}

func ConvertStructToXml(in interface{}, out *interface{}) (err error) {
	*out, err = xml.Marshal(in)
	return err
}

func ConvertJsonToStruct(in interface{}, out *interface{}) (err error) {
	return json.Unmarshal(in.([]byte), out)
}

func ConvertStructToJson(in interface{}, out *interface{}) (err error) {
	*out, err = json.Marshal(in)
	return err
}

func ConvertStructToToml(in interface{}, out *interface{}) (err error) {
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
