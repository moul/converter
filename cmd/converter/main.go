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

	app.Before = hookBefore

	app.Commands = []cli.Command{}
	for _, filter := range RegisteredConverters {
		command := cli.Command{
			Name:         filter.Name,
			Usage:        fmt.Sprintf("%s  ->  %s", filter.InputType, filter.OutputType),
			Action:       Action,
			BashComplete: BashComplete,
		}
		app.Commands = append(app.Commands, command)
	}

	app.Run(os.Args)
}

func BashComplete(c *cli.Context) {
	for _, filter := range RegisteredConverters {
		fmt.Println(filter.Name)
	}
}

func hookBefore(c *cli.Context) error {
	// configure logrus
	return nil
}

func Action(c *cli.Context) {
	args := append([]string{c.Command.Name}, c.Args()...)
	if len(args) == 0 {
		logrus.Fatalf("You need to use at least one filter")
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
