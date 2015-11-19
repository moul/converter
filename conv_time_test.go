package converter

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func ExampleConvertUnixDateToTime() {
	input := `Sat Mar  7 11:06:39.1234 PST 2015`
	var output interface{}
	ConvertUnixDateToTime(input, &output)
	fmt.Printf("%+v\n", output)
	// Output: 2015-03-07 11:06:39.1234 +0000 PST
}

func ExampleConvertANSICToTime() {
	input := `Sat Mar  7 11:06:39.1234 2015`
	var output interface{}
	ConvertANSICToTime(input, &output)
	fmt.Printf("%+v\n", output)
	// Output: 2015-03-07 11:06:39.1234 +0000 UTC
}

func ExampleConvertRFC822ToTime() {
	input := `07 Mar 15 11:06 MST`
	var output interface{}
	ConvertRFC822ToTime(input, &output)
	fmt.Printf("%+v\n", output)
	// Output: 2015-03-07 11:06:00 +0000 MST
}

func ExampleConvertRFC850ToTime() {
	input := `Saturday, 07-Mar-15 11:06:39 MST`
	var output interface{}
	ConvertRFC850ToTime(input, &output)
	fmt.Printf("%+v\n", output)
	// Output: 2015-03-07 11:06:39 +0000 MST
}

func ExampleConvertRFC1123ToTime() {
	input := `Sat, 07 Mar 2015 11:06:39 MST`
	var output interface{}
	ConvertRFC1123ToTime(input, &output)
	fmt.Printf("%+v\n", output)
	// Output: 2015-03-07 11:06:39 +0000 MST
}

/* fixme: broken test
func ExampleConvertRFC3339ToTime() {
}
*/

func TestConvertUnixDateToTime(t *testing.T) {
	Convey("Testing ConvertUnixDateToTime", t, func() {
		input := `Sat Mar  7 11:06:39.1234 PST 2015`
		var output interface{}
		err := ConvertUnixDateToTime(input, &output)
		So(err, ShouldBeNil)
		year, month, day := output.(time.Time).Date()
		zone, offset := output.(time.Time).Zone()
		So(year, ShouldEqual, 2015)
		So(month, ShouldEqual, time.March)
		So(day, ShouldEqual, 7)
		So(zone, ShouldEqual, "PST")
		So(offset, ShouldEqual, 0)
		So(output.(time.Time).Hour(), ShouldEqual, 11)
		So(output.(time.Time).Minute(), ShouldEqual, 6)
		So(output.(time.Time).Second(), ShouldEqual, 39)
		So(output.(time.Time).Nanosecond(), ShouldEqual, 123400000)
	})
}

func ExampleConvertTimeToUnix() {
	input := time.Date(1988, time.January, 25, 13, 10, 42, 0, time.UTC)
	var output interface{}
	ConvertTimeToUnix(input, &output)
	fmt.Printf("%+v\n", output.(int64))
	// Output: 570114642
}
