package app

import (
	"flag"
	"fmt"
	"github.com/pfouilloux/protoc-gen-bob/internal/display"
)

// App represents the application
type App interface {
	// Run starts the app and returns it's exit code
	Run(args []string) int
}

type cli struct {
	display display.Display
}

// Cli creates a new command line App
func Cli(display display.Display) App {
	return &cli{display}
}

// Run implements App.Run
func (cli *cli) Run(args []string) int {
	flag.Usage = func() { cli.display.Inform(cli.usage()) }

	if len(args) > 0 {
		flag.CommandLine.Parse(args[1:])
	}

	if flag.NArg() < 1 {
		cli.display.Alert(fmt.Sprintf("missing path\n%s", cli.usage()))
		return 1
	}

	return 0
}

func (cli cli) usage() string {
	return fmt.Sprintln("Usage of BobGen: bob <path> [opts]")
}
