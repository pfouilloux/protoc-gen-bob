package main

import (
	"github.com/pfouilloux/protoc-gen-bob/internal/app"
	"github.com/pfouilloux/protoc-gen-bob/internal/display"
	"os"
)

func main() {
	os.Exit(app.Cli(display.New(os.Stdout, os.Stderr)).Run(os.Args))
}
