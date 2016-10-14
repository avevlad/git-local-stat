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
	rex := hh.DateRangeFilter("", "Sun, 2 Oct 2016 18:08:09 +0300")
	fmt.Println(len(rex))
	fmt.Println(len(hh))
	fmt.Println("AFTER X", time.Since(xd))

	//
	//xd2 := time.Now()
	//fmt.Println("BEFORE X2", time.Since(start))
	//X(hh, "Sun, 2 Oct 2016 18:08:09 +0300")
	//fmt.Println("AFTER X2", time.Since(xd2))
	//fmt.Print(hh[0].VerificationFlag)

	fmt.Println("TSSS")
	fmt.Println("//////////\\///313/23////d/f/vc/c/vc/xv/cx/vd/fas/df/ds/fa")
	fmt.Println("//////////\\///313/23////d/f/vc/c/vc/xv/cx/vd/fas/df/ds/fa")
	fmt.Println("//////////\\///313/23////d/f/vc/c/vc/xv/cx/vd/fas/df/ds/fa")
	fmt.Println("//////////\\///313/23////d/f/vc/c/vc/xv/cx/vd/fas/df/ds/fa")
	//fmt.Println(len(XX(hh, "Sun, 2 Oct 2016 18:08:09 +0300", "")))

	//fmt.Print(len(XX(hh, "", "")))
}

func DateFilter(commitArray []server.GitCommitResponse, dateStart string, dateEnd string) []server.GitCommitResponse {
	res := []server.GitCommitResponse{}

	timeStart, errNoStart := server.GitDateTime(dateStart)
	timeEnd, errNoEnd := server.GitDateTime(dateEnd)

	switch true {
	case errNoStart != nil && errNoEnd != nil:
		fmt.Println("CASE NO INTERVAL")
		return commitArray
	case errNoStart != nil && errNoEnd == nil:
		fmt.Println("CASE NO START")
		for _, v := range commitArray {
			commitDate, err := server.GitDateTime(v.Author.Date)
			if commitDate.Unix() <= timeEnd.Unix() && err == nil {
				//fmt.Println(k, v.Author.Date)
				res = append(res, v)
			}
		}
	case errNoStart == nil && errNoEnd != nil:
		fmt.Println("CASE NO END")
		for _, v := range commitArray {
			commitDate, err := server.GitDateTime(v.Author.Date)
			if commitDate.Unix() >= timeStart.Unix() && err == nil {
				//fmt.Println(k, v.Author.Date)
				res = append(res, v)
			}
		}
	case errNoStart == nil && errNoEnd == nil:
		fmt.Println("CASE EXACT")
		for _, v := range commitArray {
			commitDate, err := server.GitDateTime(v.Author.Date)
			if commitDate.Unix() >= timeStart.Unix() && commitDate.Unix() <= timeEnd.Unix() && err == nil {
				//fmt.Println(k, v.Author.Date)
				res = append(res, v)
			}
		}
	}

	fmt.Println("LENENELELENELN", len(res))
	return res
}