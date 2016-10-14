package server

import "fmt"

type Predicate func(a GitCommitResponse) bool

func (commitsSlice GitCommitResponseSlice) Filter(f []Predicate) GitCommitResponseSlice {
	res := GitCommitResponseSlice{}
	for _, v := range commitsSlice {
		isSatisfied := true
		for _, ff := range f {
			if !ff(v) {
				isSatisfied = false
			}
		}

		if isSatisfied {
			res = append(res, v)
		}
	}
	return res
}

func (commitArray GitCommitResponseSlice) DateRangeFilter(dateStart string, dateEnd string, f []Predicate) GitCommitResponseSlice {
	res := GitCommitResponseSlice{}
	timeStart, errNoStart := GitDateTime(dateStart)
	timeEnd, errNoEnd := GitDateTime(dateEnd)
	switch true {
	case errNoStart != nil && errNoEnd != nil:
		fmt.Println("CASE NO INTERVAL")
		return commitArray
	case errNoStart != nil && errNoEnd == nil:
		fmt.Println("CASE NO START")
		res = commitArray.Filter(append(f, func(commit GitCommitResponse) bool {
			commitTime, err := GitDateTime(commit.Author.Date)
			if err == nil {
				return commitTime.Unix() <= timeEnd.Unix()
			}
			return false
		},))
	case errNoStart == nil && errNoEnd != nil:
		fmt.Println("CASE NO END")
		res = commitArray.Filter([]Predicate{func(commit GitCommitResponse) bool {
			commitTime, err := GitDateTime(commit.Author.Date)
			if err == nil {
				return commitTime.Unix() >= timeStart.Unix()
			}
			return false
		},})
	case errNoStart == nil && errNoEnd == nil:
		fmt.Println("CASE EXACT")
		res = commitArray.Filter([]Predicate{func(commit GitCommitResponse) bool {
			commitTime, err := GitDateTime(commit.Author.Date)
			if err == nil {
				return commitTime.Unix() >= timeStart.Unix() && commitTime.Unix() <= timeEnd.Unix()
			}
			return false
		},})
	}
	fmt.Println("LENENELELENELN", len(res))
	return res
}