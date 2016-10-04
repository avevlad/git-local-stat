package server

import (
	"testing"
	"fmt"
)


//func main() {
//}


func TestGitNormalize(t *testing.T) {
	var plainResponse = `{
  "CommitId": "b00806e054fe716a64abb48c26943585f9f747b8",
  "TreeId": "019f6c003cfa30ea258acef923767a8345c8c04e",
 "ParentId": "1eb2d88c9ee526af80011328ef3d40ff2526e34d",
  "Subject": "dsadasdsa",
  "Body": <GSXwj2snXa>
  q "qqqq<GSXwj2snXa>,
 "VerificationFlag": "N",
  "Author": {
    "Name": "AveVlad",
    "Email": "dev@vld.me",
    "Date": "Wed, 28 Sep 2016 04:00:41 +0300"
  }
  },
{
  "CommitId": "1eb2d88c9ee526af80011328ef3d40ff2526e34d",
  "TreeId": "c1d1f784c86a6f1c62974e7a5b47435324219d27",
 "ParentId": "",
  "Subject": "commit-1",
  "Body": <GSXwj2snXa>SUPER

DESC

""

1231
<GSXwj2snXa>,
 "VerificationFlag": "N",
  "Author": {
    "Name": "AveVlad",
    "Email": "dev@vld.me",
    "Date": "Wed, 28 Sep 2016 03:51:59 +0300"
  }
  },`

	var normalizeResponse = NormalizeGitGetResponse(plainResponse, "<GSXwj2snXa>")
	fmt.Println(len(normalizeResponse))
	//fmt.Println(normalizeResponse)
}
