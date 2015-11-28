package converter

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func init() {
	RegisterConverter(NewConverter("md5").SetTypes("[]byte", "string").SetConversionFunc(ConvertBytesToMd5sum))
	RegisterConverter(NewConverter("sha1").SetTypes("[]byte", "string").SetConversionFunc(ConvertBytesToSha1sum))
}

func ConvertBytesToMd5sum(in interface{}, out *interface{}) (err error) {
	*out = fmt.Sprintf("%x", md5.Sum(in.([]byte)))
	return err
}

func ConvertBytesToSha1sum(in interface{}, out *interface{}) (err error) {
	*out = fmt.Sprintf("%x", sha1.Sum(in.([]byte)))
	return err
}
