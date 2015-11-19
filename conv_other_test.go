package converter

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

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

func ExampleFetchUrlToBytes() {
	input := "http://sapin-as-a-service.appspot.com/?size=3"
	var output interface{}
	FetchUrlToBytes(input, &output)
	fmt.Printf("%+s\n", string(output.([]byte)))
	// Output:
	//           *
	//          ***
	//         *****
	//        *******
	//         *****
	//        *******
	//       *********
	//      ***********
	//     *************
	//      ***********
	//     *************
	//    ***************
	//   *****************
	//  *******************
	// *********************
	//          |||
	//          |||
	//          |||
}
