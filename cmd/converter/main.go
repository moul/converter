package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"moul.io/converter"
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
	for _, filter := range converter.RegisteredConverters {
		command := cli.Command{
			Name:         filter.Name,
			Usage:        fmt.Sprintf("%s  ->  %s", filter.InputType, filter.OutputType),
			Action:       Action,
			BashComplete: BashComplete,
		}
		app.Commands = append(app.Commands, command)
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatalf("run error: %v", err)
	}
}

func BashComplete(c *cli.Context) {
	for _, filter := range converter.RegisteredConverters {
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

	flow, err := converter.NewFlow(args)
	if err != nil {
		logrus.Fatalf("Failed to create a converter: %v", err)
	}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		logrus.Fatalf("Failed to read from stdin: %v", err)
	}

	conversionFunc, err := flow.ConversionFunc("[]byte", "interface{}")
	if err != nil {
		logrus.Fatalf("Failed to generate a conversion func: %v", err)
	}

	var output interface{}
	if err = conversionFunc(input, &output); err != nil {
		logrus.Fatalf("Failed to convert: %v", err)
	}

	fmt.Printf("%v\n", output)
}
