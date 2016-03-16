package util

import (
	"time"
)

func GetDate(date time.Time) string {
	
	
//	year, month, day := date.Date()
//	
//	newDate := strconv.Itoa(year)+"/"+strconv.Itoa(int(month))+"/"+strconv.Itoa(day)
	
	const layout = "01/02/2006"
	
	newDate := date.Format(layout)
	return newDate
}
