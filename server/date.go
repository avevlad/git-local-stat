package server

import (
	"time"
	"strconv"
)

func GitDateTime(gitDate string) (timeObj time.Time, err error) {
	layout := GitDateTimeFormat
	//timeObj, err = time.Parse(layout, gitDate)
	return time.Parse(layout, gitDate)
}

func GetDate(date time.Time) Day {
	year, month, day := date.Date()
	dateStruct := Day{}
	dateStruct.Year = year
	dateStruct.Month = month
	dateStruct.Day = day
	return dateStruct
}

func GetDateString(date time.Time) string {
	year, month, day := date.Date()
	return strconv.Itoa(year) + month.String() + strconv.Itoa(day)
}