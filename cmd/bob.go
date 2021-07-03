package main

import (
	"gitlab.com/pfouilloux/bobgen/internal/app"
	"gitlab.com/pfouilloux/bobgen/internal/display"
	"os"
)

func main() {
	os.Exit(app.Cli(display.New(os.Stdout, os.Stderr)).Run(os.Args))
}
