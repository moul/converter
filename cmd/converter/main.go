package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Sirupsen/logrus"
	. "github.com/moul/converter"
)

func main() {
	if len(os.Args) < 2 {
		logrus.Fatalf("Usage: './converter --list-filters' or './converter <filter> [filter...]'")
	}

	if os.Args[1] == "--list-filters" {
		fmt.Println("Available filters:")
		for _, filter := range RegisteredConverters {
			fmt.Printf("- %s\n", filter.Name)
		}
		return
	}

	chain, err := NewConverterChain(os.Args[1:])
	if err != nil {
		logrus.Fatalf("Failed to create a converter: %v", err)
	}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		logrus.Fatalf("Failed to read from stdin: %v", err)
	}

	conversionFunc, err := chain.ConversionFunc("[]byte", "interface{}")
	if err != nil {
		logrus.Fatalf("Failed to generate a conversion func: %v", err)
	}

	var output interface{}
	if err = conversionFunc(input, &output); err != nil {
		logrus.Fatalf("Failed to convert: %v", err)
	}

	fmt.Printf("%v\n", output)
}
