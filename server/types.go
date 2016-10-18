package server

import "time"

type GitAuthorResponse struct {
	Date  string `json:"Date"`
	Email string `json:"Email"`
	Name  string `json:"Name"`
}

type GitCommitResponse struct {
	CommitID             string `json:"CommitId"`
	TreeId               string `json:"TreeId"`
	ParentId             string `json:"ParentId"`
	Subject              string `json:"Subject"`
	SanitizedSubjectLine string `json:"SanitizedSubjectLine"`
	Body                 string `json:"Body"`
	VerificationFlag     string `json:"VerificationFlag"`
	Author               GitAuthorResponse `json:"Author"`
}

type GitCommitResponseSlice []GitCommitResponse

type GitCommitResponsesByDay struct {
	Date Day
	Commits GitCommitResponseSlice
	Count int
}

type Day struct {
	Year int
	Month time.Month
	Day int
}