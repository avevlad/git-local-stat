package server

type GitAuthorResponse struct {
	Name  string
	Email string
	Date  string
}

type GitCommitResponse struct {
	CommitId             string
	TreeId               string
	ParentId             string
	Subject              string
	SanitizedSubjectLine string
	Body                 string
	VerificationFlag     string
	Author            GitAuthorResponse
}
