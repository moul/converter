package converter_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"moul.io/converter"
)

func TestConverter(t *testing.T) {
	cases := []struct {
		input             interface{}
		converters        []string
		expectedOutput    interface{}
		expectedRunError  bool
		expectedInitError bool
	}{
		{"42", []string{"_string-to-int64"}, int64(42), false, false},
		{"HELLO WORLD", []string{"lower"}, "hello world", false, false},
		{"hello world", []string{"no-exists"}, "", false, true},
		{"hello world", []string{"rev"}, "dlrow olleh", false, false},
		{"hello world", []string{"title"}, "Hello World", false, false},
		{"hello world", []string{"upper"}, "HELLO WORLD", false, false},
		{"hello world", []string{}, "hello world", false, false},
		{"hello world", nil, "hello world", false, false},
		{"42", []string{"_string-to-float64"}, float64(42), false, false},
		{"42.42", []string{"_string-to-int64"}, nil, true, false},
		{"42.42", []string{"_string-to-float64"}, float64(42.42), false, false},
		{int64(42), []string{"_int64-to-string"}, "42", false, false},
		{float64(42), []string{"_float64-to-string"}, "42", false, false},
		{float64(42.5), []string{"_float64-to-string"}, "42.5", false, false},
		{[]byte("hello world"), []string{"_bytes-to-string"}, "hello world", false, false},
		{"hello world", []string{"_parse-string"}, "hello world", false, false},
		{"42", []string{"_parse-string"}, int64(42), false, false},
		{"42.42", []string{"_parse-string"}, float64(42.42), false, false},
		{"hello world", []string{"rev", "rev"}, "hello world", false, false},
		{"hello world", []string{"rev", "upper"}, "DLROW OLLEH", false, false},
		{"hello world", []string{"upper", "rev"}, "DLROW OLLEH", false, false},
		{"hello world", []string{"_string-to-bytes", "hex"}, "68656c6c6f20776f726c64", false, false},
		{"hello world", []string{"_string-to-bytes", "hex", "hex-decode", "_bytes-to-string"}, "hello world", false, false},
		{"hello world", []string{"_string-to-bytes", "md5", "hex"}, "5eb63bbbe01eeed093cb22bb8f5acdc3", false, false},
		{"hello world", []string{"_string-to-bytes", "md5", "md5", "hex"}, "241d8a27c836427bd7f04461b60e7359", false, false},
		{"hello world", []string{"_string-to-bytes", "sha1", "hex"}, "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed", false, false},
		{"hello world", []string{"_string-to-bytes", "base32"}, "NBSWY3DPEB3W64TMMQ======", false, false},
		{"hello world", []string{"_string-to-bytes", "hexbase32"}, "D1IMOR3F41RMUSJCCG======", false, false},
		{"hello world", []string{"_string-to-bytes", "base58"}, "StV1DL6CwTryKyV", false, false},
		{"hello world", []string{"_string-to-bytes", "base64"}, "aGVsbG8gd29ybGQ=", false, false},
		{"hello world", []string{"_string-to-bytes", "urlbase64"}, "aGVsbG8gd29ybGQ=", false, false},
		{"hello world", []string{"_string-to-bytes", "rawurlbase64"}, "aGVsbG8gd29ybGQ", false, false},
		{"hello world", []string{"_string-to-bytes", "base32", "base32-decode", "_bytes-to-string"}, "hello world", false, false},
		{"hello world", []string{"_string-to-bytes", "hexbase32", "hexbase32-decode", "_bytes-to-string"}, "hello world", false, false},
		{"hello world", []string{"_string-to-bytes", "base58", "base58-decode", "_bytes-to-string"}, "hello world", false, false},
		{"hello world", []string{"_string-to-bytes", "base64", "base64-decode", "_bytes-to-string"}, "hello world", false, false},
		{"hello world", []string{"_string-to-bytes", "urlbase64", "urlbase64-decode", "_bytes-to-string"}, "hello world", false, false},
		{"hello world", []string{"_string-to-bytes", "rawurlbase64", "rawurlbase64-decode", "_bytes-to-string"}, "hello world", false, false},
	}

	for _, tc := range cases {
		name := fmt.Sprintf("%v_%s", tc.input, strings.Join(tc.converters, ","))
		t.Run(name, func(t *testing.T) {
			fn, err := converter.ChainFunc(tc.converters)
			if tc.expectedInitError {
				require.Error(t, err)
				require.Nil(t, fn)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, fn)

			ret, err := fn(tc.input)
			if tc.expectedRunError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedOutput, ret)
			}
		})
	}
}

func Example() {
	ret, _ := converter.Chain("hello world", []string{"rev", "upper"})
	fmt.Println(ret)
	// Output: DLROW OLLEH
}
