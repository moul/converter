package converter

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConvertBytesToString(t *testing.T) {
	Convey("Testing ConvertBytesToString", t, func() {
		input := []byte("hello world!")
		var output interface{}
		err := ConvertBytesToString(input, &output)
		So(err, ShouldBeNil)
		So(output, ShouldEqual, "hello world!")
	})
}

func ExampleConvertBytesToString() {
	var output interface{}
	ConvertBytesToString([]byte("hello world!"), &output)
	fmt.Println(output)
	// Output: hello world!
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

func ExampleConvertStringToBytes() {
	var output interface{}
	ConvertStringToBytes("hello world!", &output)
	fmt.Println(output)
	// Output: [104 101 108 108 111 32 119 111 114 108 100 33]
}

func TestConvertIntegerToString(t *testing.T) {
	Convey("Testing ConvertIntegerToString", t, func() {
		input := int64(1234567890)
		var output interface{}
		err := ConvertIntegerToString(input, &output)
		So(err, ShouldBeNil)
		So(output, ShouldEqual, "1234567890")
	})
}

func ExampleConvertIntegerToString() {
	var output interface{}
	ConvertIntegerToString(int64(1234567890), &output)
	fmt.Println(output)
	// Output: 1234567890
}

func TestConvertStringToInteger(t *testing.T) {
	Convey("Testing ConvertStringToInteger", t, func() {
		input := "1234567890"
		var output interface{}
		err := ConvertStringToInteger(input, &output)
		So(err, ShouldBeNil)
		So(output, ShouldEqual, 1234567890)
	})
}

func ExampleConvertStringToInteger() {
	var output interface{}
	ConvertStringToInteger("1234567890", &output)
	fmt.Println(output)
	// Output: 1234567890
}

func TestConvertStringToFloat(t *testing.T) {
	Convey("Testing ConvertStringToFloat", t, func() {
		input := "3.1415"
		var output interface{}
		err := ConvertStringToFloat(input, &output)
		So(err, ShouldBeNil)
		So(output, ShouldEqual, 3.1415)
	})
}

func ExampleConvertStringToFloat() {
	var output interface{}
	ConvertStringToFloat("3.1415", &output)
	fmt.Println(output)
	// Output: 3.1415
}

func TestConvertFloatToString(t *testing.T) {
	Convey("Testing ConvertStringToFloat", t, func() {
		input := 3.1415
		var output interface{}
		err := ConvertFloatToString(input, &output)
		So(err, ShouldBeNil)
		So(output, ShouldEqual, "3.1415")
	})
}

func ExampleConvertFloatToString() {
	var output interface{}
	ConvertFloatToString(3.1415, &output)
	fmt.Println(output)
	// Output: 3.1415
}
