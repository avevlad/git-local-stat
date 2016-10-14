package server

import "fmt"

func (commitArray GitCommitResponseSlice) DateRangeFilter(dateStart string, dateEnd string) GitCommitResponseSlice {
	res := GitCommitResponseSlice{}

	timeStart, errNoStart := GitDateTime(dateStart)
	timeEnd, errNoEnd := GitDateTime(dateEnd)

	switch true {
	case errNoStart != nil && errNoEnd != nil:
		fmt.Println("CASE NO INTERVAL")
		return commitArray
	case errNoStart != nil && errNoEnd == nil:
		fmt.Println("CASE NO START")
		for _, v := range commitArray {
			commitDate, err := GitDateTime(v.Author.Date)
			if commitDate.Unix() <= timeEnd.Unix() && err == nil {
				//fmt.Println(k, v.Author.Date)
				res = append(res, v)
			}
		}
	case errNoStart == nil && errNoEnd != nil:
		fmt.Println("CASE NO END")
		for _, v := range commitArray {
			commitDate, err := GitDateTime(v.Author.Date)
			if commitDate.Unix() >= timeStart.Unix() && err == nil {
				//fmt.Println(k, v.Author.Date)
				res = append(res, v)
			}
		}
	case errNoStart == nil && errNoEnd == nil:
		fmt.Println("CASE EXACT")
		for _, v := range commitArray {
			commitDate, err := GitDateTime(v.Author.Date)
			if commitDate.Unix() >= timeStart.Unix() && commitDate.Unix() <= timeEnd.Unix() && err == nil {
				//fmt.Println(k, v.Author.Date)
				res = append(res, v)
			}
		}
	}

	fmt.Println("LENENELELENELN", len(res))
	return res
}