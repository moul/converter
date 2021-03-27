package converter

import (
	"bytes"
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

	"github.com/BurntSushi/toml"
	"github.com/mr-tron/base58"
)

// Converters is a map containing converters that can be called using a name.
var Converters = map[string]interface{}{
	"base32":              base32Encode,
	"base32-decode":       base32Decode,
	"base58":              base58Encode,
	"base58-decode":       base58Decode,
	"base64":              base64Encode,
	"base64-decode":       base64Decode,
	"hex":                 hexEncode,
	"hex-decode":          hexDecode,
	"hexbase32":           hexbase32Encode,
	"hexbase32-decode":    hexbase32Decode,
	"lower":               lower,
	"md5":                 md5Sum,
	"rawurlbase64":        rawurlbase64Encode,
	"rawurlbase64-decode": rawurlbase64Decode,
	"rev":                 rev,
	"sha1":                sha1Sum,
	"title":               title,
	"upper":               upper,
	"urlbase64":           urlbase64Encode,
	"urlbase64-decode":    urlbase64Decode,
	"json":                jsonMarshal,
	"toml":                tomlEncode,
	"json-decode":         jsonDecode,
	"csv-decode":          csvDecode,
	"xml":                 xmlEncode,
	"xml-decode":          xmlDecode,

	// internal
	"_parse-string":      parseString,
	"_bytes-to-string":   bytesToString,
	"_string-to-bytes":   stringToBytes,
	"_int64-to-string":   int64ToString,
	"_string-to-int64":   stringToInt64,
	"_string-to-float64": stringToFloat64,
	"_float64-to-string": float64ToString,
	/*
	   fetch                 string  ->  []byte
	   sleep-100ms           interface{}  ->  interface{}
	   sleep-1s              interface{}  ->  interface{}
	   sleep-2s              interface{}  ->  interface{}
	   sleep-5s              interface{}  ->  interface{}
	   sleep-10s             interface{}  ->  interface{}
	   sleep-1m              interface{}  ->  interface{}
	   reverse               string  ->  string
	   upper                 string  ->  string
	   lower                 string  ->  string
	   split-lines           []byte  ->  []byte
	   to-unix               time.Time  ->  int64
	   parse-ansi-date       string  ->  time.Time
	   parse-rfc339-date     string  ->  time.Time
	   parse-rfc822-date     string  ->  time.Time
	   parse-rfc850-date     string  ->  time.Time
	   parse-rfc1123-date    string  ->  time.Time
	   parse-unix-date       string  ->  time.Time
	   parse-date            string  ->  time.Time
	   time-to-string        time.Time  ->  string
	   parse-unix-timestamp  int64  ->  time.Time
	*/
}

// parseString takes a string in input and tries to cast it in a more specific type (date, int, etc).
// This function should be the first one to be called in a chain when using a CLI.
func parseString(in string) interface{} {
	if n, err := strconv.ParseInt(in, 10, 0); err == nil {
		return n
	}
	if n, err := strconv.ParseFloat(in, 64); err == nil {
		return n
	}

	// FIXME: try to parse other formats first
	return in
}

func rev(in string) (string, error) {
	runes := []rune(in)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes), nil
}

func tomlEncode(in interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := toml.NewEncoder(buf).Encode(in)
	return buf.Bytes(), err
}

func jsonDecode(in []byte) (interface{}, error) {
	var out interface{}
	err := json.Unmarshal(in, &out)
	return out, err
}

func csvDecode(in string) (interface{}, error) {
	r := csv.NewReader(strings.NewReader(in))
	return r.ReadAll()
}

func xmlEncode(in interface{}) ([]byte, error) {
	return xml.Marshal(in)
}

func xmlDecode(in []byte) (interface{}, error) {
	var out interface{}
	err := xml.Unmarshal(in, &out)
	return out, err
}

func base32Encode(in []byte) string {
	return base32.StdEncoding.EncodeToString(in)
}

func base32Decode(in string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(in)
}

func base58Encode(in []byte) string {
	return base58.Encode(in)
}

func base58Decode(in string) ([]byte, error) {
	return base58.Decode(in)
}

func base64Encode(in []byte) string {
	return base64.StdEncoding.EncodeToString(in)
}

func base64Decode(in string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(in)
}

func urlbase64Encode(in []byte) string {
	return base64.URLEncoding.EncodeToString(in)
}

func urlbase64Decode(in string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(in)
}

func rawurlbase64Encode(in []byte) string {
	return base64.RawURLEncoding.EncodeToString(in)
}

func rawurlbase64Decode(in string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(in)
}

func hexEncode(in []byte) string {
	return fmt.Sprintf("%x", in)
}

func hexDecode(in string) ([]byte, error) {
	return hex.DecodeString(in)
}

func hexbase32Encode(in []byte) string {
	return base32.HexEncoding.EncodeToString(in)
}

func hexbase32Decode(in string) ([]byte, error) {
	return base32.HexEncoding.DecodeString(in)
}

func md5Sum(in []byte) []byte {
	ret := md5.Sum(in)
	return ret[:]
}

func sha1Sum(in []byte) []byte {
	ret := sha1.Sum(in)
	return ret[:]
}

func lower(in string) string {
	return strings.ToLower(in)
}

func upper(in string) string {
	return strings.ToUpper(in)
}

func title(in string) string {
	return strings.Title(in)
}

func jsonMarshal(in interface{}) ([]byte, error) {
	return json.Marshal(in)
}

func bytesToString(in []byte) string {
	return string(in)
}

func stringToBytes(in string) []byte {
	return []byte(in)
}

func int64ToString(in int64) string {
	return strconv.FormatInt(in, 10)
}

func stringToInt64(in string) (int64, error) {
	return strconv.ParseInt(strings.TrimSpace(in), 10, 0)
}

func stringToFloat64(in string) (float64, error) {
	return strconv.ParseFloat(strings.TrimSpace(in), 64)
}

func float64ToString(in float64) string {
	return strconv.FormatFloat(in, 'f', -1, 64)
}
