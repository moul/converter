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

func TestConvertBytesToBase32(t *testing.T) {
	Convey("Testing ConvertBytesToBase32", t, func() {
		input := []byte("hello world!")
		var output interface{}
		err := ConvertBytesToBase32(input, &output)
		So(err, ShouldBeNil)
		So(output, ShouldEqual, "NBSWY3DPEB3W64TMMQQQ====")
	})
}

func TestConvertBase32ToBytes(t *testing.T) {
	Convey("Testing ConvertBase32ToBytes", t, func() {
		input := "NBSWY3DPEB3W64TMMQQQ===="
		var output interface{}
		err := ConvertBase32ToBytes(input, &output)
		So(err, ShouldBeNil)
		So(output, ShouldResemble, []byte("hello world!"))
	})
}

func TestConvertBytesToString(t *testing.T) {
	Convey("Testing ConvertBytesToString", t, func() {
		input := []byte("hello world!")
		var output interface{}
		err := ConvertBytesToString(input, &output)
		So(err, ShouldBeNil)
		So(output, ShouldEqual, "hello world!")
	})
}

func TestConvertStringToBytes(t *testing.T) {
	Convey("Testing ConvertStringToBytes", t, func() {
		input := "hello world!"
		var output interface{}
		err := ConvertStringToBytes(input, &output)
		So(err, ShouldBeNil)
		So(output, ShouldResemble, []byte("hello world!"))
	})
}
