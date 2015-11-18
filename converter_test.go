package converter

import (
	"fmt"
	"testing"
	"time"

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
	})
}

func TestConvertBytesToBase64(t *testing.T) {
	Convey("Testing ConvertBytesToBase64", t, func() {
		input := []byte("hello world!")
		var output interface{}
		err := ConvertBytesToBase64(input, &output)
		So(err, ShouldBeNil)
		So(output, ShouldEqual, "aGVsbG8gd29ybGQh")
	})
}

func ExampleConvertBytesToBase64() {
	var output interface{}
	ConvertBytesToBase64([]byte("hello world!"), &output)
	fmt.Println(output)
	// Output: aGVsbG8gd29ybGQh
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

func ExampleConvertBase64ToBytes() {
	var output interface{}
	ConvertBase64ToBytes("aGVsbG8gd29ybGQh", &output)
	fmt.Println(output)
	// Output: [104 101 108 108 111 32 119 111 114 108 100 33]
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

func ExampleConvertBytesToBase32() {
	var output interface{}
	ConvertBytesToBase32([]byte("hello world!"), &output)
	fmt.Println(output)
	// Output: NBSWY3DPEB3W64TMMQQQ====
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

func ExampleConvertBase32ToBytes() {
	var output interface{}
	ConvertBase32ToBytes("NBSWY3DPEB3W64TMMQQQ====", &output)
	fmt.Println(output)
	// Output: [104 101 108 108 111 32 119 111 114 108 100 33]
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

func TestConvertBytesToHex(t *testing.T) {
	Convey("Testing ConvertBytesToHex", t, func() {
		input := []byte("hello world!")
		var output interface{}
		err := ConvertBytesToHex(input, &output)
		So(err, ShouldBeNil)
		So(output, ShouldEqual, "68656c6c6f20776f726c6421")
	})
}

func ExampleConvertBytesToHex() {
	var output interface{}
	ConvertBytesToHex([]byte("hello world!"), &output)
	fmt.Println(output)
	// Output: 68656c6c6f20776f726c6421
}

func TestConvertHexToBytes(t *testing.T) {
	Convey("Testing ConvertHexToBytes", t, func() {
		input := "68656c6c6f20776f726c6421"
		var output interface{}
		err := ConvertHexToBytes(input, &output)
		So(err, ShouldBeNil)
		So(output, ShouldResemble, []byte("hello world!"))
	})
}

func ExampleConvertHexToBytes() {
	var output interface{}
	ConvertHexToBytes("68656c6c6f20776f726c6421", &output)
	fmt.Println(output)
	// Output: [104 101 108 108 111 32 119 111 114 108 100 33]
}

func TestConvertIntegerToString(t *testing.T) {
	Convey("Testing ConvertIntegerToString", t, func() {
		input := 1234567890
		var output interface{}
		err := ConvertIntegerToString(input, &output)
		So(err, ShouldBeNil)
		So(output, ShouldEqual, "1234567890")
	})
}

func ExampleConvertIntegerToString() {
	var output interface{}
	ConvertIntegerToString(1234567890, &output)
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

func ExampleConvertJsonToStruct() {
	var output interface{}
	input := []byte(`["Hello",42,3.1415]`)
	ConvertJsonToStruct(input, &output)
	fmt.Printf("%+v\n", output)
	// Output: [Hello 42 3.1415]
}

func TestConvertJsonToStruct(t *testing.T) {
	Convey("Testing ConvertJsonToStruct", t, func() {
		input := []byte(`["Hello",42,3.1415]`)
		var output interface{}
		err := ConvertJsonToStruct(input, &output)
		So(err, ShouldBeNil)
		So(output.([]interface{})[0], ShouldEqual, "Hello")
		So(output.([]interface{})[1], ShouldEqual, int(42))
		So(output.([]interface{})[2], ShouldEqual, 3.1415)
	})
}

func ExampleConvertStructToJson() {
	var output interface{}
	input := []interface{}{
		"Hello", 42, 3.1415,
	}
	ConvertStructToJson(input, &output)
	fmt.Printf("%s\n", output)
	// Output: ["Hello",42,3.1415]
}

func TestConvertStructToJson(t *testing.T) {
	Convey("Testing ConvertStructToJson", t, func() {
		input := []interface{}{
			"Hello", 42, 3.1415,
		}
		var output interface{}
		err := ConvertStructToJson(input, &output)
		So(err, ShouldBeNil)
		So(output, ShouldResemble, []byte(`["Hello",42,3.1415]`))
	})
}

func ExampleConvertStringToCsv() {
	input := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
"Manfred",Touron,moul
`
	var output interface{}
	ConvertStringToCsv(input, &output)
	fmt.Printf("%+v\n", output)
	// Output: [[first_name last_name username] [Rob Pike rob] [Ken Thompson ken] [Robert Griesemer gri] [Manfred Touron moul]]
}

func TestConvertStringToCsv(t *testing.T) {
	Convey("Testing ConvertStringToCsv", t, func() {
		input := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
"Manfred",Touron,moul
`
		var output interface{}
		err := ConvertStringToCsv(input, &output)
		So(err, ShouldBeNil)
		So(output, ShouldResemble, [][]string{
			[]string{"first_name", "last_name", "username"},
			[]string{"Rob", "Pike", "rob"},
			[]string{"Ken", "Thompson", "ken"},
			[]string{"Robert", "Griesemer", "gri"},
			[]string{"Manfred", "Touron", "moul"},
		})
	})
}

func TestDelayer(t *testing.T) {
	Convey("Testing Delayer", t, func() {
		duration := time.Millisecond * 100
		delayerFunc := Delayer(time.Millisecond * 100)

		Convey("Testing on bytes", func() {
			input := []byte("hello world!")
			var output interface{}
			t1 := time.Now()
			err := delayerFunc(input, &output)
			t2 := time.Now()
			So(err, ShouldBeNil)
			So(output, ShouldResemble, input)
			So(t2.Sub(t1) >= duration, ShouldBeTrue)
		})
		Convey("Testing on integer", func() {
			input := 1234567890
			var output interface{}
			t1 := time.Now()
			err := delayerFunc(input, &output)
			t2 := time.Now()
			So(err, ShouldBeNil)
			So(output, ShouldEqual, input)
			So(t2.Sub(t1) >= duration, ShouldBeTrue)
		})
		Convey("Testing on float", func() {
			input := 3.1415
			var output interface{}
			t1 := time.Now()
			err := delayerFunc(input, &output)
			t2 := time.Now()
			So(err, ShouldBeNil)
			So(output, ShouldEqual, input)
			So(t2.Sub(t1) >= duration, ShouldBeTrue)
		})
	})
}
