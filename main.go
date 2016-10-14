package main

import (
	"os"
	"github.com/jessevdk/go-flags"
	"github.com/avevlad/git-local-stat/server"
	"fmt"
	"time"
)

var options server.CommandLineOptions
var parser = flags.NewParser(&options, flags.Default)

func main() {
	start := time.Now()

	if _, err := parser.Parse(); err != nil {
		os.Exit(1)
	}
	server.Init(options)
	bodyCommits := server.GetBodyCommits()
	//_ = server.SetBody(&server.Commits[0], bodyCommits[server.Commits[0].CommitId])

	hh := server.SetBodies(server.Commits, bodyCommits)
	//fmt.Print("hhhhhhhhhhhhhh", hh[0].Author.Date, hh[0].CommitID)
	xd := time.Now()
	fmt.Println("BEFORE X", time.Since(start))
	//X(hh, "Sun, 1 Oct 2016 18:08:09 +0300")
	rex := hh.DateRangeFilter("", "", []server.Predicate{
		func(commit server.GitCommitResponse) bool {
			return commit.Author.Name == "AveVlad"
		},
		func(commit server.GitCommitResponse) bool {
			return commit.CommitID == "24d4f30ce4001a2ab9a37005a1d2c6537ecd3ac0"
		}}...)
	//rexx := hh.DateRangeFilter("Sun, 2 Oct 2016 18:08:09 +0300", "")
	fmt.Println(len(rex))
	//fmt.Println(len(rexx))
	fmt.Println(len(hh))
	fmt.Println("AFTER X", time.Since(xd))
	fmt.Println("//////////\\///313/23////d/f/vc/c/vc/xv/cx/vd/fas/df/ds/fa")
	fmt.Println("//////////\\///313/23////d/f/vc/c/vc/xv/cx/vd/fas/df/ds/fa")
	fmt.Println("//////////\\///313/23////d/f/vc/c/vc/xv/cx/vd/fas/df/ds/fa")
	fmt.Println("//////////\\///313/23////d/f/vc/c/vc/xv/cx/vd/fas/df/ds/fa")
	//fmt.Println(len(XX(hh, "Sun, 2 Oct 2016 18:08:09 +0300", "")))

	//fmt.Print(len(XX(hh, "", "")))
}