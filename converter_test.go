package converter

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_manual_chaining(t *testing.T) {
	Convey("Testing chaining", t, func() {
		input := "hello world!"
		var output1 interface{}
		var err error

		err = ConvertStringToBytes(input, &output1)
		So(err, ShouldBeNil)
		So(output1, ShouldResemble, []byte("hello world!"))

		var output2 interface{}
		err = ConvertBytesToHex(output1, &output2)
		So(err, ShouldBeNil)
		So(output2, ShouldEqual, "68656c6c6f20776f726c6421")

		var output3 interface{}
		err = ConvertHexToBytes(output2, &output3)
		So(err, ShouldBeNil)
		So(output3, ShouldResemble, []byte("hello world!"))

		var output4 interface{}
		err = ConvertBytesToString(output3, &output4)
		So(err, ShouldBeNil)
		So(output4, ShouldEqual, input)
	})
}

func TestPipe(t *testing.T) {
	Convey("Testing Pipe", t, func() {
		Convey(`string("hello world!") | ConvertStringToBytes | ConvertBytesToBase64`, func() {
			input := "hello world!"
			var output interface{}

			pipeFunc := Pipe(ConvertStringToBytes, ConvertBytesToBase64)
			err := pipeFunc(input, &output)
			So(err, ShouldBeNil)
			So(output, ShouldEqual, "aGVsbG8gd29ybGQh")
		})
	})
}

func TestChain(t *testing.T) {
	Convey("Testing Chain", t, func() {
		Convey(`float64(3.1415) | ConvertFloatToString | ConvertStringToBytes | ConvertBytesToBase64 | ConvertStringToBytes | ConvertBytesToBase32`, func() {
			input := 3.1415
			var output interface{}

			chainFunc := Chain(ConvertFloatToString, ConvertStringToBytes, ConvertBytesToBase64, ConvertStringToBytes, ConvertBytesToBase32)

			err := chainFunc(input, &output)
			So(err, ShouldBeNil)
			So(output, ShouldEqual, "JV4TI6COIRCTC===")
		})
		Convey(``, func() {
			input := "http://httpbin.org/headers"
			var output interface{}

			chainFunc := Chain(FetchUrlToBytes, ConvertJsonToStruct, ConvertStructToToml, ConvertBytesToString)

			err := chainFunc(input, &output)
			So(err, ShouldBeNil)
			So(output, ShouldEqual, `[headers]
  Accept-Encoding = "gzip"
  Host = "httpbin.org"
  User-Agent = "Go-http-client/1.1"
`)
		})
	})
}

func TestConversionToStreamConv(t *testing.T) {
	Convey("Testing ConversionToStreamConv", t, func() {
		streamFn := ConversionToStreamConv(ConvertStringToBytes)
		in := make(chan interface{}, 10)
		out := streamFn(in)

		in <- "hello world!"
		in <- "HELLO WORLD!"
		in <- "HeLlO wOrLd!"
		So(<-out, ShouldResemble, []byte("hello world!"))
		So(<-out, ShouldResemble, []byte("HELLO WORLD!"))
		So(<-out, ShouldResemble, []byte("HeLlO wOrLd!"))
		// select {
		// case _, ok := <-out:
		// 	So(ok, ShouldBeFalse)
		// }
	})
}

func TestStreamPipe(t *testing.T) {
	Convey("Testing StreamPipe", t, func() {
		streamPipe := StreamPipe(ConversionToStreamConv(ConvertStringToBytes), ConversionToStreamConv(ConvertBytesToBase64))
		in := make(chan interface{}, 10)
		out := streamPipe(in)

		in <- "hello world!"
		in <- "HELLO WORLD!"
		in <- "HeLlO wOrLd!"
		So(<-out, ShouldEqual, "aGVsbG8gd29ybGQh")
		So(<-out, ShouldEqual, "SEVMTE8gV09STEQh")
		So(<-out, ShouldEqual, "SGVMbE8gd09yTGQh")
		// select {
		// case _, ok := <-out:
		// 	So(ok, ShouldBeFalse)
		// }
	})
}

func TestStreamChain(t *testing.T) {
	Convey("Testing StreamChain", t, func() {
		streamChain := StreamChain(ConversionToStreamConv(ConvertFloatToString), ConversionToStreamConv(ConvertStringToBytes), ConversionToStreamConv(ConvertBytesToBase64), ConversionToStreamConv(ConvertStringToBytes), ConversionToStreamConv(ConvertBytesToBase32))
		in := make(chan interface{}, 10)
		out := streamChain(in)

		in <- 3.1415
		in <- 4242.4242
		in <- -123456789.987654321
		So(<-out, ShouldEqual, "JV4TI6COIRCTC===")
		So(<-out, ShouldEqual, "JZCESMCNNE2DATLKKF4Q====")
		So(<-out, ShouldEqual, "JRKEK6KNPJITCTTKMM2E6UZUGVHUIYZSJZKFC6SNO46T2===")
		// select {
		// case _, ok := <-out:
		// 	So(ok, ShouldBeFalse)
		// }
	})
}
