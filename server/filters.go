package server

import "fmt"

type Predicate func(commit GitCommitResponse) bool

// If strict = true, Filter All
// Else Filter Any
func (commitsSlice GitCommitResponseSlice) Filter(strict bool, filters ...Predicate) GitCommitResponseSlice {
	res := GitCommitResponseSlice{}
	for _, v := range commitsSlice {
		isSatisfied := strict
		for _, filter := range filters {
			f := filter(v)
			if strict {
				f = !f
			}
			if f {
				isSatisfied = !strict
				break
			}
		}

		if isSatisfied {
			res = append(res, v)
		}
	}
	return res
}

func (commitArray GitCommitResponseSlice) DateRangeFilter(dateStart string, dateEnd string, strict bool, f ...Predicate) GitCommitResponseSlice {
	res := GitCommitResponseSlice{}
	timeStart, errNoStart := GitDateTime(dateStart)
	timeEnd, errNoEnd := GitDateTime(dateEnd)
	switch true {
	case errNoStart != nil && errNoEnd != nil:
		fmt.Println("CASE NO INTERVAL")
		if f != nil {
			res = commitArray.Filter(strict, f...)
		} else {
			return commitArray
		}
	case errNoStart != nil && errNoEnd == nil:
		fmt.Println("CASE NO START")
		res = commitArray.Filter(strict, append(f, func(commit GitCommitResponse) bool {
			commitTime, err := GitDateTime(commit.Author.Date)
			if err == nil {
				return commitTime.Unix() <= timeEnd.Unix()
			}
			return false
		})...)
	case errNoStart == nil && errNoEnd != nil:
		fmt.Println("CASE NO END")
		res = commitArray.Filter(strict, append(f, func(commit GitCommitResponse) bool {
			commitTime, err := GitDateTime(commit.Author.Date)
			if err == nil {
				return commitTime.Unix() >= timeStart.Unix()
			}
			return false
		})...)
	case errNoStart == nil && errNoEnd == nil:
		fmt.Println("CASE EXACT")
		res = commitArray.Filter(strict, append(f, func(commit GitCommitResponse) bool {
			commitTime, err := GitDateTime(commit.Author.Date)
			if err == nil {
				return commitTime.Unix() >= timeStart.Unix() && commitTime.Unix() <= timeEnd.Unix()
			}
			return false
		})...)
	}
	fmt.Println("LENENELELENELN", len(res))
	return res
}