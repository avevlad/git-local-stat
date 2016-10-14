package server

import (
	"testing"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"strconv"
)

func TestJsonParseCommit(t *testing.T) {
	str := "{\n\"CommitId\": \"75897c2dcd1dd3a6ca46284dd37e13d22b4b16b4\",\n\"TreeId\": \"d9585ffa48a5563525251e8cc657c61d1a3f4625\",\n\"ParentId\": \"\",\n\"Subject\": \"Initial public release\",\n\"SanitizedSubjectLine\": \"Initial-public-release\",\n\"Body\": \"\",\n\"VerificationFlag\": \"N\",\n\"Author\": {\n\"Name\": \"Paul O’Shannessy\",\n\"Email\": \"paul@oshannessy.com\",\n\"Date\": \"Wed, 29 May 2013 12:46:11 -0700\"\n}\n}"
	res := GitCommitResponse{}
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		t.Error(err)
	}
}

func TestJsonParseCommitList(t *testing.T) {
	str := "[{\n\"CommitId\": \"75897c2dcd1dd3a6ca46284dd37e13d22b4b16b4\",\n\"TreeId\": \"d9585ffa48a5563525251e8cc657c61d1a3f4625\",\n\"ParentId\": \"\",\n\"Subject\": \"Initial public release\",\n\"SanitizedSubjectLine\": \"Initial-public-release\",\n\"Body\": \"\",\n\"VerificationFlag\": \"N\",\n\"Author\": {\n\"Name\": \"Paul O’Shannessy\",\n\"Email\": \"paul@oshannessy.com\",\n\"Date\": \"Wed, 29 May 2013 12:46:11 -0700\"\n}\n}]"
	res := GitCommitResponseSlice{}
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		t.Error(err)
	}
}

func TestJsonParseInvalidCharacterCommitList(t *testing.T) {
	//re := regexp.MustCompile(`(?U)<div>(.*)<div>`)
	re := regexp.MustCompile(fmt.Sprint(`(?U)`, `<div>`, `(.*)`, `<div>`))
	//
	input := `11<div>qq""\||32(&)*&%^$&^#q<div>22, 55<div>r5<div>77`
	//
	//fmt.Println(re.FindAllString(input, -1))
	//fmt.Println(len(re.FindAllString(input, -1)))

	//match, _ := regexp.Compile()
	//fmt.Println(match)

	for _, element := range re.FindAllString(input, -1) {
		fmt.Println(element)
		fmt.Println(strings.Replace(element, "div", "11", -1))
	}

	fmt.Println(strconv.Quote(`1"2'3`))

}