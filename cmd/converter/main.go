package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	. "github.com/moul/converter"
)

var (
	VERSION   string
	GITCOMMIT string
)

func main() {
	app := cli.NewApp()
	app.Name = "converter"
	app.Author = "Manfred Touron"
	app.Email = "https://github.com/moul/converter"
	app.Version = VERSION + " (" + GITCOMMIT + ")"
	app.EnableBashCompletion = true
	app.BashComplete = BashComplete

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "list-filters",
			Usage: "List available filters",
		},
	}

	app.Before = hookBefore
	app.Action = Action

	app.Run(os.Args)
}

func BashComplete(c *cli.Context) {
	if len(c.Args()) == 0 {
		fmt.Println("--list-filters")
	}
	for _, filter := range RegisteredConverters {
		fmt.Println(filter.Name)
	}
}

func hookBefore(c *cli.Context) error {
	// configure logrus
	return nil
}

func Action(c *cli.Context) {
	if c.Bool("list-filters") {
		fmt.Println("Available filters:")
		for _, filter := range RegisteredConverters {
			fmt.Printf("- %s\n", filter.Name)
		}
		return
	}

	args := c.Args()
	if len(args) == 0 {
		logrus.Fatalf("You need to use at least one filter")
	}

	for _, arg := range args {
		if arg == "--generate-bash-completion" {
			return
		}
	}

	chain, err := NewConverterChain(args)
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
