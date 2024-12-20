package main

import (
	"fmt"

	"github.com/awee-ai/cli"
)

type InitVars struct {
	Test      bool `arg:"test" help:"Test flag"`
	SomeStruc struct {
		Verbose bool `arg:"v" help:"Verbose output"`
	} `arg:"some-struc" help:"Some struct"`
}

type InitCommand struct {
	cli.BaseCommand[InitVars]
}

var _ cli.Command[InitVars] = (*InitCommand)(nil)

func NewInitCommand() *InitCommand {
	return &InitCommand{BaseCommand: cli.NewBaseCommand[InitVars]()}
}

func (c *InitCommand) Run(options cli.GlobalOptions, unknowns cli.Unknowns) error {
	fmt.Println("Init command")
	return nil
}

func (c *InitCommand) Validate(vars map[string]any) error {
	return nil
}

func (c *InitCommand) Help() string {
	return "Display help"
}
