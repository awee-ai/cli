package main

import (
	"fmt"

	"github.com/contentforward/cli"
)

type ExampleVars struct {
	Verbose bool `arg:"-v" help:"Verbose output"`
}

type ExampleCommand struct {
	cli.BaseCommand[ExampleVars]
}

var _ cli.Command[ExampleVars] = (*ExampleCommand)(nil)

func NewExampleCommand() *ExampleCommand {
	return &ExampleCommand{BaseCommand: cli.NewBaseCommand[ExampleVars]()}
}

func (c *ExampleCommand) Run(options cli.GlobalOptions) error {
	fmt.Println("Example command")
	return nil
}

func (c *ExampleCommand) Validate(vars map[string]any) error {
	return nil
}

func (c *ExampleCommand) Help() string {
	return "Example command"
}