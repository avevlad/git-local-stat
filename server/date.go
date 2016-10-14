package server

import (
	"time"
)

func GitDateTime(gitDate string) (timeObj time.Time, err error) {
	layout := GitDateTimeFormat
	//timeObj, err = time.Parse(layout, gitDate)
	return time.Parse(layout, gitDate)
}
