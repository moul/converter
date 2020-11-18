package converter

import (

	// nolint:gosec
	"crypto/md5"

	// nolint:gosec
	"crypto/sha1"
	"fmt"
)

// nolint:gochecknoinits // need a refactor to remove it
func init() {
	RegisterConverter(NewConverter("md5").SetTypes("[]byte", "string").SetConversionFunc(ConvertBytesToMd5sum))
	RegisterConverter(NewConverter("sha1").SetTypes("[]byte", "string").SetConversionFunc(ConvertBytesToSha1sum))
}

func ConvertBytesToMd5sum(in interface{}, out *interface{}) (err error) {
	*out = fmt.Sprintf("%x", md5.Sum(in.([]byte))) // nolint:gosec
	return err
}

func ConvertBytesToSha1sum(in interface{}, out *interface{}) (err error) {
	*out = fmt.Sprintf("%x", sha1.Sum(in.([]byte))) // nolint:gosec // expected
	return err
}
