package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"moul.io/converter"
)

func main() {
	convertor := converter.Chain(converter.ConvertJSONToStruct, converter.ConvertStructToTOML)
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		logrus.Fatalf("Failed to read from stdin: %v", err)
	}
	var output interface{}
	if err = convertor(input, &output); err != nil {
		logrus.Fatalf("Failed to convert from json to toml: %v", err)
	}
	fmt.Printf("%s", output.([]byte))
}
