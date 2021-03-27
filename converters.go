package converter

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/mr-tron/base58"
)

// Converters is a map containing converters that can be called using a name.
var Converters = map[string]interface{}{
	"base32":              base32.StdEncoding.EncodeToString,
	"base32-decode":       base32.StdEncoding.DecodeString,
	"base58":              base58.Encode,
	"base58-decode":       base58.Decode,
	"base64":              base64.StdEncoding.EncodeToString,
	"base64-decode":       base64.StdEncoding.DecodeString,
	"hex":                 func(in []byte) string { return fmt.Sprintf("%x", in) },
	"hex-decode":          hex.DecodeString,
	"hexbase32":           base32.HexEncoding.EncodeToString,
	"hexbase32-decode":    base32.HexEncoding.DecodeString,
	"lower":               strings.ToLower,
	"md5":                 func(in []byte) []byte { ret := md5.Sum(in); return ret[:] },
	"rawurlbase64":        base64.RawURLEncoding.EncodeToString,
	"rawurlbase64-decode": base64.RawURLEncoding.DecodeString,
	"rev":                 rev,
	"sha1":                func(in []byte) []byte { ret := sha1.Sum(in); return ret[:] },
	"title":               strings.Title,
	"upper":               strings.ToUpper,
	"urlbase64":           base64.URLEncoding.EncodeToString,
	"urlbase64-decode":    base64.URLEncoding.DecodeString,

	// internal
	"_parse-string":      parseString,
	"_bytes-to-string":   func(in []byte) string { return string(in) },
	"_string-to-bytes":   func(in string) []byte { return []byte(in) },
	"_int64-to-string":   func(in int64) string { return strconv.FormatInt(in, 10) },
	"_string-to-int64":   func(in string) (int64, error) { return strconv.ParseInt(strings.TrimSpace(in), 10, 0) },
	"_string-to-float64": func(in string) (float64, error) { return strconv.ParseFloat(strings.TrimSpace(in), 64) },
	"_float64-to-string": func(in float64) string { return strconv.FormatFloat(in, 'f', -1, 64) },
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
