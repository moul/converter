package converter

import (
	"io/ioutil"
	"net/http"
	"time"
)

func init() {
	RegisterConverter(NewConverter("fetch").SetTypes("string", "[]byte").SetConversionFunc(FetchUrlToBytes))
	RegisterConverter(NewConverter("sleep-100ms").SetConversionFunc(HundredMillisecondDelayer))
	RegisterConverter(NewConverter("sleep-1s").SetConversionFunc(OneSecondDelayer))
	RegisterConverter(NewConverter("sleep-2s").SetConversionFunc(TwoSecondDelayer))
	RegisterConverter(NewConverter("sleep-5s").SetConversionFunc(FiveSecondDelayer))
	RegisterConverter(NewConverter("sleep-10s").SetConversionFunc(TenSecondDelayer))
	RegisterConverter(NewConverter("sleep-1m").SetConversionFunc(OneMinuteDelayer))
}

func Delayer(duration time.Duration) ConversionFn {
	return func(in interface{}, out *interface{}) error {
		time.Sleep(duration)
		*out = in
		return nil
	}
}

var HundredMillisecondDelayer = Delayer(100 * time.Millisecond)
var OneSecondDelayer = Delayer(time.Second)
var TwoSecondDelayer = Delayer(2 * time.Second)
var FiveSecondDelayer = Delayer(5 * time.Second)
var TenSecondDelayer = Delayer(10 * time.Second)
var OneMinuteDelayer = Delayer(time.Minute)

func FetchUrlToBytes(in interface{}, out *interface{}) error {
	resp, err := http.Get(in.(string))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	*out, err = ioutil.ReadAll(resp.Body)
	return err
}
