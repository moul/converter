package converter

import (
	"fmt"
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

func ExampleConvertStructToToml() {
	var output interface{}
	input := map[string]interface{}{
		"a": "Hello",
		"b": 42,
		"c": 3.1415,
	}
	ConvertStructToToml(input, &output)
	fmt.Printf("%s\n", output)
	// Output:
	// a = "Hello"
	// b = 42
	// c = 3.1415
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

/* fixme: broken test
func ExampleConvertXmlToStruct() {
	var output interface{}
	input := []byte(`<string>Hello</string><int>42</int><float64>3.1415</float64>`)
	ConvertXmlToStruct(input, &output)
	fmt.Printf("%+v\n", output)
	// Output: [Hello 42 3.1415]
}
*/

/* fixme: broken test
func TestConvertXmlToStruct(t *testing.T) {
	Convey("Testing ConvertXmlToStruct", t, func() {
		input := []byte(`<string>Hello</string><int>42</int><float64>3.1415</float64>`)
		var output interface{}
		err := ConvertXmlToStruct(input, &output)
		So(err, ShouldBeNil)
		So(output.([]interface{})[0], ShouldEqual, "Hello")
		So(output.([]interface{})[1], ShouldEqual, int(42))
		So(output.([]interface{})[2], ShouldEqual, 3.1415)
	})
}
*/

func ExampleConvertStructToXml() {
	var output interface{}
	input := []interface{}{
		"Hello", 42, 3.1415,
	}
	ConvertStructToXml(input, &output)
	fmt.Printf("%s\n", output)
	// Output: <string>Hello</string><int>42</int><float64>3.1415</float64>
}

func TestConvertStructToXml(t *testing.T) {
	Convey("Testing ConvertStructToXml", t, func() {
		input := []interface{}{
			"Hello", 42, 3.1415,
		}
		var output interface{}
		err := ConvertStructToXml(input, &output)
		So(err, ShouldBeNil)
		So(output, ShouldResemble, []byte(`<string>Hello</string><int>42</int><float64>3.1415</float64>`))
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
