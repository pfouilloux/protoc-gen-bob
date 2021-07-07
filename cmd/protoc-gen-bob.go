package main

import (
	"github.com/pfouilloux/protoc-gen-bob/internal/adapters/iostreams"
	"os"
)

func main() {
	_ = iostreams.New(os.Stdin, os.Stdout).Handle()
}
