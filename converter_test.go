package converter

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConvertBytesToBase64(t *testing.T) {
	Convey("Testing ConvertBytesToBase64", t, func() {
		input := []byte("hello world!")
		var output interface{}
		err := ConvertBytesToBase64(input, &output)
		So(err, ShouldBeNil)
		So(output, ShouldEqual, "aGVsbG8gd29ybGQh")
	})
}

func TestConvertBase64ToBytes(t *testing.T) {
	Convey("Testing ConvertBase64ToBytes", t, func() {
		input := "aGVsbG8gd29ybGQh"
		var output interface{}
		err := ConvertBase64ToBytes(input, &output)
		So(err, ShouldBeNil)
		So(output, ShouldResemble, []byte("hello world!"))
	})
}
