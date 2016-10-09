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
	fmt.Print("hhhhhhhhhhhhhh", hh[0].Author.Date, hh[0].CommitID)
	xd := time.Now()
	fmt.Println("BEFORE X", time.Since(start))
	X(hh, "Sun, 1 Oct 2016 18:08:09 +0300")
	fmt.Println("AFTER X", time.Since(xd))

	xd2 := time.Now()
	fmt.Println("BEFORE X2", time.Since(start))
	X(hh, "Sun, 2 Oct 2016 18:08:09 +0300")
	fmt.Println("AFTER X2", time.Since(xd2))
	fmt.Print(hh[0].VerificationFlag)
}

func X(commitArray []server.GitCommitResponse, date string) {
	time, _ := server.GitDateTime(date)
	for _, v := range commitArray {
		commitDate, _ := server.GitDateTime(v.Author.Date)
		if commitDate.Day() == time.Day() {
			//fmt.Println(v.Subject)
		}
	}
}