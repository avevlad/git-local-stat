package server

func (commits GitCommitResponseSlice) GroupByDay() []GitCommitResponsesByDay {
	res := []GitCommitResponsesByDay{}
	indexArray := map[string]int{}
	count := 0

	for _, commit := range commits {
		commitDate, err := GitDateTime(commit.Author.Date)
		commitDateString := GetDateString(commitDate)

		if err == nil {
			index, ok := indexArray[commitDateString]
			currentCommits := []GitCommitResponse{commit, }
			dayCommits := GitCommitResponsesByDay{}
			dayCommits.Date = GetDate(commitDate)
			dayCommits.Commits = currentCommits
			dayCommits.Count = 1

			if ok {
				// Key exists: add current commit
				currentCommits = append(res[index].Commits, commit)
				dayCommits.Commits = currentCommits
				dayCommits.Count = len(currentCommits)
				res[index] = dayCommits
			} else {
				// Key does not exist, create commit slice and append current commit
				indexArray[commitDateString] = count
				res = append(res, dayCommits)
				count++
			}
		}
	}

	return res
}

func (commits GitCommitResponseSlice) GroupByDayFromDates(dates ...string) []GitCommitResponsesByDay {
	res := []GitCommitResponsesByDay{}
	indexArray := map[string]int{}
	count := 0

	for _, commit := range commits {
		for _, date := range dates {
			commitDate, err := GitDateTime(commit.Author.Date)
			commitDateString := GetDateString(commitDate)
			dateTime, err := GitDateTime(date)
			dateTimeString := GetDateString(dateTime)

			if err == nil {
				if commitDateString == dateTimeString {
					index, ok := indexArray[commitDateString]
					currentCommits := []GitCommitResponse{commit, }
					dayCommits := GitCommitResponsesByDay{}
					dayCommits.Date = GetDate(commitDate)
					dayCommits.Commits = currentCommits
					dayCommits.Count = 1

					if ok {
						// Key exists: add current commit
						currentCommits = append(res[index].Commits, commit)
						dayCommits.Commits = currentCommits
						dayCommits.Count = len(currentCommits)
						res[index] = dayCommits
					} else {
						// Key does not exist, create commit slice and append current commit
						indexArray[commitDateString] = count
						res = append(res, dayCommits)
						count++
					}
				}
			}
		}
	}

	return res
}

func (commits GitCommitResponseSlice) GroupByWeek() []GitCommitResponsesByWeek {
	res := []GitCommitResponsesByWeek{}
	indexArray := map[string][]int{}
	count := 0

	for _, commit := range commits {
		commitDate, err := GitDateTime(commit.Author.Date)
		year, week := commitDate.ISOWeek()
		weekAndYear := string(year) + string(week)
		if err == nil {
			yearWeekIndexTuple, ok := indexArray[weekAndYear]
			currentCommits := []GitCommitResponse{commit, }
			weekCommits := GitCommitResponsesByWeek{}
			weekCommits.Commits = currentCommits
			weekCommits.Count = 1

			if ok {
				// Key exists: add current commit
				weekCommits.Year = yearWeekIndexTuple[0]
				weekCommits.Week = yearWeekIndexTuple[1]
				index := yearWeekIndexTuple[2]
				currentCommits = append(res[index].Commits, commit)
				weekCommits.Commits = currentCommits
				weekCommits.Count = len(currentCommits)
				res[index] = weekCommits
			} else {
				// Key does not exist, create commit slice and append current commit
				weekCommits.Year = year
				weekCommits.Week = week
				indexArray[weekAndYear] = []int{year, week, count}
				res = append(res, weekCommits)
				count++
			}
		}
	}

	return res
}

func (commits GitCommitResponseSlice) GroupByMonth() []GitCommitResponsesByMonth {
	res := []GitCommitResponsesByMonth{}
	indexArray := map[string][]int{}
	count := 0

	for _, commit := range commits {
		commitDate, err := GitDateTime(commit.Author.Date)
		month := int(commitDate.Month())
		year := commitDate.Year()
		yearAndMonth := string(year) + string(month)
		if err == nil {
			yearMonthIndexTuple, ok := indexArray[yearAndMonth]
			currentCommits := []GitCommitResponse{commit, }
			monthCommits := GitCommitResponsesByMonth{}
			monthCommits.Commits = currentCommits
			monthCommits.Count = 1

			if ok {
				// Key exists: add current commit
				monthCommits.Year = yearMonthIndexTuple[0]
				monthCommits.Month = yearMonthIndexTuple[1]
				index := yearMonthIndexTuple[2]
				currentCommits = append(res[index].Commits, commit)
				monthCommits.Commits = currentCommits
				monthCommits.Count = len(currentCommits)
				res[index] = monthCommits
			} else {
				// Key does not exist, create commit slice and append current commit
				monthCommits.Year = year
				monthCommits.Month = month
				indexArray[yearAndMonth] = []int{year, month, count}
				res = append(res, monthCommits)
				count++
			}
		}
	}

	return res
}

func (commits GitCommitResponseSlice) GroupByYear() []GitCommitResponsesByYear {
	res := []GitCommitResponsesByYear{}
	indexArray := map[int][]int{}
	count := 0

	for _, commit := range commits {
		commitDate, err := GitDateTime(commit.Author.Date)
		year := commitDate.Year()
		if err == nil {
			yearIndexTuple, ok := indexArray[year]
			currentCommits := []GitCommitResponse{commit, }
			monthCommits := GitCommitResponsesByYear{}
			monthCommits.Commits = currentCommits
			monthCommits.Count = 1

			if ok {
				// Key exists: add current commit
				monthCommits.Year = yearIndexTuple[0]
				index := yearIndexTuple[1]
				currentCommits = append(res[index].Commits, commit)
				monthCommits.Commits = currentCommits
				monthCommits.Count = len(currentCommits)
				res[index] = monthCommits
			} else {
				// Key does not exist, create commit slice and append current commit
				monthCommits.Year = year
				indexArray[year] = []int{year, count}
				res = append(res, monthCommits)
				count++
			}
		}
	}

	return res
}