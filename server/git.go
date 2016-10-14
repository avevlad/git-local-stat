package server

import (
	"os/exec"
	"bytes"
	"fmt"
	"os"
	"strings"
	"encoding/json"
	"strconv"
	"regexp"
)

var SanitizedKeyBodyPattern string = "<" + srand(10) + ">"

//var GitCmdPrettyFlag string = fmt.Sprint(
//	`--pretty=format:{%n  "CommitId": "%H",%n  "TreeId": "%T",%n "ParentId": "%P",%n  "Subject": "%f",%n  "Body": `,
//	SanitizedKeyBodyPattern,
//	`%b`,
//	SanitizedKeyBodyPattern,
//	`,%n "VerificationFlag": "%G?",%n  "Author": {%n    "Name": "%aN",%n    "Email": "%aE",%n    "Date": "%aD"%n  }%n  },`,

var GitCmdPrettyFlag string = fmt.Sprint(
	// unix time stamp?
	// commiter?
	// '%N': commit notes
	`--pretty=format:{%n  "CommitId": "%H",%n  "TreeId": "%T",%n "ParentId": "%P",%n  "Subject": "%f", %n "VerificationFlag": "%G?",%n  "Author": {%n    "Name": "%aN",%n    "Email": "%aE",%n    "Date": "%aD"%n  }%n  },`,
)

var GitCmdBodyFlag string = fmt.Sprint(
	`--pretty=format:%n----------%n%H$$$$$%b`,
)

func ParseGitNormalizeResponse(str string) GitCommitResponseSlice {
	res := GitCommitResponseSlice{}

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

func NormalizeGitGetBodyResponse(str string, regexKey string) (resp map[string]string) {
	var arr = strings.Split(str, "----------")
	resp = make(map[string]string)

	for index, element := range arr {
		if index == 0 && element == "\n" {
			continue
		}
		lineArr := strings.Split(element, "$$$$$")
		if (len(lineArr) == 2) {
			body := lineArr[1]
			commitId := strings.TrimSpace(lineArr[0])
			resp[commitId] = strconv.Quote(body)
		}
		//var newString = strings.Replace(element, regexKey, "", -1)
		//fmt.Println("newString", newString)
		//fmt.Println("newString", newString)
		//fmt.Println(strings.Replace(element, SanitizedBodyPatter, "11", 2))
		//res = strings.Replace(res, element, newString, -1)
	}
	return resp
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

func GitGetBodyResponse() string {
	cmd := exec.Command("git", "--no-pager", "log", GitCmdBodyFlag)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error1: %v\n", NotGitRepositoryErrorMsg)
		os.Exit(1)
	}

	return out.String()
}

func GetBodyCommits() map[string]string {
	plainBodyResponse := GitGetBodyResponse()
	normalizeResponse := NormalizeGitGetBodyResponse(plainBodyResponse, SanitizedKeyBodyPattern)
	//fmt.Print(normalizeResponse)
	return normalizeResponse
}

func GetCommits() GitCommitResponseSlice {
	plainResponse := GitGetResponse()
	//fmt.Print(plainResponse)
	normalizeResponse := NormalizeGitGetResponse(plainResponse, SanitizedKeyBodyPattern)
	return ParseGitNormalizeResponse(normalizeResponse)
}

func SetBody(resp *GitCommitResponse, body string) GitCommitResponse {
	// Mutate
	resp.Body = body
	return *resp
}

func SetBodies(in GitCommitResponseSlice, bodies map[string]string) (GitCommitResponseSlice) {
	// Mutate
	for index, elem := range in {
		if val, ok := bodies[elem.CommitID]; ok {
			in[index] = SetBody(&elem, val)
		}
	}

	return in
}

var Commits GitCommitResponseSlice = GetCommits()
