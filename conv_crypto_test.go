package converter

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConvertBytesToMd5sum(t *testing.T) {
	Convey("Testing ConvertBytesToMd5sum", t, func() {
		input := []byte("hello world!")
		var output interface{}
		err := ConvertBytesToMd5sum(input, &output)
		So(err, ShouldBeNil)
		So(output, ShouldEqual, "fc3ff98e8c6a0d3087d515c0473f8677")
	})
}

func ExampleConvertBytesToMd5sum() {
	var output interface{}
	ConvertBytesToMd5sum([]byte("hello world!"), &output)
	fmt.Println(output)
	// Output: fc3ff98e8c6a0d3087d515c0473f8677
}
func TestConvertBytesToSha1sum(t *testing.T) {
	Convey("Testing ConvertBytesToSha1sum", t, func() {
		input := []byte("hello world!")
		var output interface{}
		err := ConvertBytesToSha1sum(input, &output)
		So(err, ShouldBeNil)
		So(output, ShouldEqual, "430ce34d020724ed75a196dfc2ad67c77772d169")
	})
}

func ExampleConvertBytesToSha1sum() {
	var output interface{}
	ConvertBytesToSha1sum([]byte("hello world!"), &output)
	fmt.Println(output)
	// Output: 430ce34d020724ed75a196dfc2ad67c77772d169
}
