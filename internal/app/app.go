package app

import (
	"BobGen/internal/output"
	"flag"
	"fmt"
)

type App interface {
	Run(args []string) int
}

type cli struct {
	out output.Output
}

func Cli(out output.Output) App {
	return &cli{out}
}

func (cli *cli) Run(args []string) int {
	flag.Usage = func() { cli.out.Inform(cli.Usage()) }

	if len(args) > 0 {
		flag.CommandLine.Parse(args[1:])
	}

	if flag.NArg() != 1 {
		cli.out.Alert(fmt.Sprintf("missing path\n%s", cli.Usage()))
		return 1
	}

	return 0
}

func (cli cli) Usage() string {
	return fmt.Sprintln("Usage of BobGen: bob <path> [opts]")
}
