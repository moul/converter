package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"moul.io/converter"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run() error {
	fn, err := converter.ChainFunc([]string{"json-decode", "toml", "_bytes-to-string"})
	if err != nil {
		return err
	}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	ret, err := fn(input)
	if err != nil {
		return err
	}

	fmt.Println(ret)
	return nil
}
