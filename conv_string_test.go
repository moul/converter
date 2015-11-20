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

func TestStreamBufferSplitLines(t *testing.T) {
	Convey("Testing StreamBufferSplitLines", t, func() {
		in := make(chan interface{}, 10)
		out := StreamBufferSplitLines(in)

		in <- []byte("hello world\nwhat's up ?\nok bye.")
		in <- []byte("where is bryan ?\nbryan is in the kitchen.")
		So(<-out, ShouldResemble, []byte("hello world"))
		So(<-out, ShouldResemble, []byte("what's up ?"))
		So(<-out, ShouldResemble, []byte("ok bye."))
		So(<-out, ShouldResemble, []byte("where is bryan ?"))
		So(<-out, ShouldResemble, []byte("bryan is in the kitchen."))
		// select {
		// case _, ok := <-out:
		// 	So(ok, ShouldBeFalse)
		// }
	})
}
