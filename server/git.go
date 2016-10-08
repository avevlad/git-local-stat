package server

import (
	"os/exec"
	"bytes"
	"fmt"
	"os"
	"strings"
	"encoding/json"
	"regexp"
)

var SanitizedKeyBodyPatter string = "<" + srand(10) + ">"

//var GitCmdPrettyFlag string = fmt.Sprint(
//	`--pretty=format:{%n  "CommitId": "%H",%n  "TreeId": "%T",%n "ParentId": "%P",%n  "Subject": "%f",%n  "Body": `,
//	SanitizedKeyBodyPatter,
//	`%b`,
//	SanitizedKeyBodyPatter,
//	`,%n "VerificationFlag": "%G?",%n  "Author": {%n    "Name": "%aN",%n    "Email": "%aE",%n    "Date": "%aD"%n  }%n  },`,

var GitCmdPrettyFlag string = fmt.Sprint(
	`--pretty=format:{%n  "CommitId": "%H",%n  "TreeId": "%T",%n "ParentId": "%P",%n  "Subject": "%f", %n "VerificationFlag": "%G?",%n  "Author": {%n    "Name": "%aN",%n    "Email": "%aE",%n    "Date": "%aD"%n  }%n  },`,
	//`--pretty=format:{%n  "CommitId": "%H",%n  "TreeId": "%T",%n "ParentId": "%P",%n  "Subject": "%f", %n  "Body": "%b", %n "VerificationFlag": "%G?",%n  "Author": {%n    "Name": "%aN",%n    "Email": "%aE",%n    "Date": "%aD"%n  }%n  },`,
)

func ParseGitNormalizeResponse(str string) []GitCommitResponse {
	res := []GitCommitResponse{}
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error:2 %v\n", err)
		os.Exit(1)
	}

	return res
}

func NormalizeGitGetResponse(str string, regexKey string) string {
	brokenDblQuoteRe := regexp.MustCompile("\".*\"\\:\\s\"(.*\".*)\"")
	matchBrokenDblQuote := brokenDblQuoteRe.FindStringSubmatch(str)
	fixedFslash := str
	if len(matchBrokenDblQuote) == 2 {
		brokenDblQuoteElem := matchBrokenDblQuote[1]
		fixedDblQuoteElem := strings.Replace(brokenDblQuoteElem, "\"", "U+201C", -1)
		fixedDblQuoteStr := strings.Replace(str, brokenDblQuoteElem, fixedDblQuoteElem, -1)
		fixedBslash := strings.Replace(fixedDblQuoteStr, "\\", "U+005C", -1)
		fixedFslash = strings.Replace(fixedBslash, "/", "U+002F", -1)
	}

	//re := regexp.MustCompile(fmt.Sprint(regexKey, `((.|\n)*?)`, regexKey))
	//var res = str
	//fmt.Println("----", len(re.FindAllString(str, -1)))
	//for index := range re.FindAllString(str, -1) {
	//	fmt.Println("!!!!!!!!!!!!", index)
	//
	//	//fmt.Println("newString", newString)
	//	newString = strconv.Quote(newString)
	//	//fmt.Println("newString", newString)
	//	res = strings.Replace(res, element, newString, -1)
	//}
	//fmt.Println("--------222")
	return "[" + strings.TrimSuffix(fixedFslash, ",") + "]"
}

func GitGetResponse() string {
	cmd := exec.Command("git", "--no-pager", "log", GitCmdPrettyFlag)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error1: %v\n", NotGitRepositoryErrorMsg)
		os.Exit(1)
	}

	return out.String()
}

func GetCommits() []GitCommitResponse {
	plainResponse := GitGetResponse()
	//fmt.Println(plainResponse)
	normalizeResponse := NormalizeGitGetResponse(plainResponse, SanitizedKeyBodyPatter)
	return ParseGitNormalizeResponse(normalizeResponse)
}

var Commits []GitCommitResponse = GetCommits()
