package converter

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func ConvertBytesToMd5sum(in interface{}, out *interface{}) (err error) {
	*out = fmt.Sprintf("%x", md5.Sum(in.([]byte)))
	return err
}

func ConvertBytesToSha1sum(in interface{}, out *interface{}) (err error) {
	*out = fmt.Sprintf("%x", sha1.Sum(in.([]byte)))
	return err
}
