package converter

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReverseString(t *testing.T) {
	Convey("Testing ReverseString", t, func() {
		Convey("Testing on bytes", func() {
			input := "Hello Worldz!"
			var output interface{}
			err := ReverseString(input, &output)
			So(err, ShouldBeNil)
			So(output, ShouldEqual, "!zdlroW olleH")
		})
	})
}

func TestUppercase(t *testing.T) {
	Convey("Testing Uppercase", t, func() {
		Convey("Testing on bytes", func() {
			input := "Hello Worldz!"
			var output interface{}
			err := Uppercase(input, &output)
			So(err, ShouldBeNil)
			So(output, ShouldEqual, "HELLO WORLDZ!")
		})
	})
}

func TestLowercase(t *testing.T) {
	Convey("Testing Lowercase", t, func() {
		Convey("Testing on bytes", func() {
			input := "Hello Worldz!"
			var output interface{}
			err := Lowercase(input, &output)
			So(err, ShouldBeNil)
			So(output, ShouldEqual, "hello worldz!")
		})
	})
}
