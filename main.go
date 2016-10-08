package main

import (
	"os"
	"github.com/jessevdk/go-flags"
	"github.com/avevlad/git-local-stat/server"
	"fmt"
)

var options server.CommandLineOptions
var parser = flags.NewParser(&options, flags.Default)

func main() {
	if _, err := parser.Parse(); err != nil {
		os.Exit(1)
	}
	server.Init(options)
	bodyCommits := server.GetBodyCommits()
	//_ = server.SetBody(&server.Commits[0], bodyCommits[server.Commits[0].CommitId])

	hh := server.SetBodies(server.Commits, bodyCommits)
	fmt.Print("hhhhhhhhhhhhhh", hh[2])
}