package main

import (
	"BobGen/internal/app"
	"BobGen/internal/output"
	"os"
)

func main() {
	os.Exit(app.Cli(output.New(os.Stdout, os.Stderr)).Run(os.Args))
}
