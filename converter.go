package converter

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base32"
	"encoding/base64"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"
)

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

func ConvertBytesToMd5sum(in interface{}, out *interface{}) (err error) {
	*out = fmt.Sprintf("%x", md5.Sum(in.([]byte)))
	return err
}

func ConvertBytesToSha1sum(in interface{}, out *interface{}) (err error) {
	*out = fmt.Sprintf("%x", sha1.Sum(in.([]byte)))
	return err
}

func ConvertJsonToStruct(in interface{}, out *interface{}) (err error) {
	return json.Unmarshal(in.([]byte), out)
}

func ConvertStructToJson(in interface{}, out *interface{}) (err error) {
	*out, err = json.Marshal(in)
	return err
}

func ConvertXmlToStruct(in interface{}, out *interface{}) (err error) {
	return xml.Unmarshal(in.([]byte), out)
}

func ConvertStructToXml(in interface{}, out *interface{}) (err error) {
	*out, err = xml.Marshal(in)
	return err
}

func ConvertStringToCsv(in interface{}, out *interface{}) (err error) {
	r := csv.NewReader(strings.NewReader(in.(string)))
	*out, err = r.ReadAll()
	return err
}

func Delayer(duration time.Duration) ConversionFunc {
	return func(in interface{}, out *interface{}) error {
		time.Sleep(duration)
		*out = in
		return nil
	}
}

var HundredMillisecondDelayer = Delayer(100 * time.Millisecond)
var OneSecondDelayer = Delayer(time.Second)
var TwoSecondDelayer = Delayer(2 * time.Second)
var FiveSecondDelayer = Delayer(5 * time.Second)
var TenSecondDelayer = Delayer(10 * time.Second)
var OneMinuteDelayer = Delayer(time.Minute)

func DateToTimeConverter(format string) ConversionFunc {
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
