package util

import (
	"time"
	"strconv"
)

func GetDate(date time.Time) string {
	
	
	year, month, day := date.Date()
	
	newDate := strconv.Itoa(year)+"/"+strconv.Itoa(int(month))+"/"+strconv.Itoa(day)
	
	return newDate
}
