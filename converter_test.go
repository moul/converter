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
		input          interface{}
		converters     []string
		expectedOutput interface{}
		expectedError  bool
	}{
		{"hello world", []string{"rev"}, "dlrow olleh", false},
		{"hello world", []string{"upper"}, "HELLO WORLD", false},
		{"HELLO WORLD", []string{"lower"}, "hello world", false},
		{"hello world", []string{"title"}, "Hello World", false},
		{"42", []string{"_string-to-int64"}, int64(42), false},
		{"42", []string{"_string-to-float64"}, float64(42), false},
		{"42.42", []string{"_string-to-int64"}, nil, true},
		{"42.42", []string{"_string-to-float64"}, float64(42.42), false},
		{int64(42), []string{"_int64-to-string"}, "42", false},
		{float64(42), []string{"_float64-to-string"}, "42", false},
		{float64(42.5), []string{"_float64-to-string"}, "42.5", false},
		{"hello world", []string{"_parse-string"}, "hello world", false},
		{"42", []string{"_parse-string"}, int64(42), false},
		{"42.42", []string{"_parse-string"}, float64(42.42), false},
		{"hello world", []string{"rev", "rev"}, "hello world", false},
		{"hello world", []string{"rev", "upper"}, "DLROW OLLEH", false},
		{"hello world", []string{"upper", "rev"}, "DLROW OLLEH", false},
		{"hello world", []string{"_string-to-bytes", "hex"}, "68656c6c6f20776f726c64", false},
		{"hello world", []string{"_string-to-bytes", "md5", "hex"}, "5eb63bbbe01eeed093cb22bb8f5acdc3", false},
		{"hello world", []string{"_string-to-bytes", "md5", "md5", "hex"}, "241d8a27c836427bd7f04461b60e7359", false},
		{"hello world", []string{"_string-to-bytes", "sha1", "hex"}, "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed", false},
	}

	for _, tc := range cases {
		name := fmt.Sprintf("%v_%s", tc.input, strings.Join(tc.converters, ","))
		t.Run(name, func(t *testing.T) {
			fn, err := converter.ChainFunc(tc.converters)
			require.NoError(t, err)
			require.NotNil(t, fn)

			ret, err := fn(tc.input)
			if tc.expectedError {
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
