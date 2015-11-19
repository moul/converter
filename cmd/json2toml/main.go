package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Sirupsen/logrus"
	. "github.com/moul/converter"
)

func main() {
	convertor := Chain(ConvertJsonToStruct, ConvertStructToToml)
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
