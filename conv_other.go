package converter

import (
	"io/ioutil"
	"net/http"
	"time"
)

// nolint:gochecknoinits // need a refactor to remove it
func init() {
	RegisterConverter(NewConverter("fetch").SetTypes("string", "[]byte").SetConversionFunc(FetchURLToBytes))
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

var (
	HundredMillisecondDelayer = Delayer(100 * time.Millisecond) // nolint:gomnd
	OneSecondDelayer          = Delayer(time.Second)            // nolint:gomnd
	TwoSecondDelayer          = Delayer(2 * time.Second)        // nolint:gomnd
	FiveSecondDelayer         = Delayer(5 * time.Second)        // nolint:gomnd
	TenSecondDelayer          = Delayer(10 * time.Second)       // nolint:gomnd
	OneMinuteDelayer          = Delayer(time.Minute)            // nolint:gomnd
)

func FetchURLToBytes(in interface{}, out *interface{}) error {
	resp, err := http.Get(in.(string)) // nolint:noctx
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	*out, err = ioutil.ReadAll(resp.Body)
	return err
}
