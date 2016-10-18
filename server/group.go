package server

func (commits GitCommitResponseSlice) GroupByDay() []GitCommitResponsesByDay {
	res := []GitCommitResponsesByDay{}
	indexArray := map[string]int{}
	dayCount := 0

	for _, commit := range commits {
		commitDate, err := GitDateTime(commit.Author.Date)

		if err == nil {
			index, ok := indexArray[GetDateString(commitDate)]
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
				indexArray[GetDateString(commitDate)] = dayCount
				res = append(res, dayCommits)
				dayCount++
			}
		}
	}

	return res
}

func (commits GitCommitResponseSlice) GroupByDayFromDates(dates ...string) []GitCommitResponsesByDay {
	res := []GitCommitResponsesByDay{}
	indexArray := map[string]int{}
	dayCount := 0

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
						indexArray[commitDateString] = dayCount
						res = append(res, dayCommits)
						dayCount++
					}
				}
			}
		}
	}

	return res
}