package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"

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
	if VERSION != "" || GITCOMMIT != "" {
		app.Version = VERSION + " (" + GITCOMMIT + ")"
	}
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{}
	filters := []string{}
	for filter := range converter.Converters {
		filters = append(filters, filter)
	}
	sort.Strings(filters)
	for _, filter := range filters {
		command := cli.Command{
			Name:         filter,
			Action:       Action,
			BashComplete: BashComplete,
			Hidden:       filter[0] == '_',
			// Usage: fmt.Sprintf("%s  ->  %s", filter.InputType, filter.OutputType),
		}
		app.Commands = append(app.Commands, command)
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func BashComplete(c *cli.Context) {
	for filter := range converter.Converters {
		fmt.Println(filter)
	}
}

func Action(c *cli.Context) error {
	args := append([]string{c.Command.Name}, c.Args()...)
	if len(args) == 0 {
		return fmt.Errorf("you need to use at least one filter")
	}

	fn, err := converter.ChainFunc(args)
	if err != nil {
		return fmt.Errorf("Failed to create a converter: %w", err)
	}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("Failed to read from stdin: %w", err)
	}

	output, err := fn(input)
	if err != nil {
		return fmt.Errorf("Failed to convert: %w", err)
	}

	fmt.Printf("%v\n", output)
	return nil
}
