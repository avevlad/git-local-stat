package main

import (
	"os"
	"github.com/jessevdk/go-flags"
	"github.com/avevlad/git-local-stat/server"
)

var options server.CommandLineOptions
var parser = flags.NewParser(&options, flags.Default)

func main() {
	if _, err := parser.Parse(); err != nil {
		os.Exit(1)
	}

	server.Init(options)
}