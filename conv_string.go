package converter

import "strings"

func ReverseString(in interface{}, out *interface{}) (err error) {
	runes := []rune(in.(string))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	*out = string(runes)
	return nil
}

func Uppercase(in interface{}, out *interface{}) (err error) {
	*out = strings.ToUpper(in.(string))
	return nil
}

func Lowercase(in interface{}, out *interface{}) (err error) {
	*out = strings.ToLower(in.(string))
	return nil
}
