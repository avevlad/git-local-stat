package server

import (
	"os/exec"
	"bytes"
	"fmt"
	"os"
	"strings"
	"encoding/json"
	"regexp"
	"strconv"
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
)

func ParseGitNormalizeResponse(str string) []GitCommitResponse {
	//fmt.Println(str)
	res := []GitCommitResponse{}
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error:2 %v\n", err)
		os.Exit(1)
	}

	return res
}

func NormalizeGitGetResponse(str string, regexKey string) string {
	re := regexp.MustCompile(fmt.Sprint(regexKey, `((.|\n)*?)`, regexKey))
	var res = str
	fmt.Println("----", len(re.FindAllString(str, -1)))
	for index, element := range re.FindAllString(str, -1) {
		//fmt.Println("element", element)
		fmt.Println("index", index)
		var newString = strings.Replace(element, regexKey, "", -1)
		//fmt.Println("newString", newString)
		newString = strconv.Quote(newString)
		//fmt.Println("newString", newString)
		//fmt.Println(strings.Replace(element, SanitizedBodyPatter, "11", 2))
		res = strings.Replace(res, element, newString, -1)
	}
	fmt.Println("--------222")
	return "[" + strings.TrimSuffix(res, ",") + "]"
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
