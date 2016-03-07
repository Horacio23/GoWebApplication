package util

import (
	"time"
	//"strconv"
	"fmt"
)

func GetDate(date time.Time) string {
	
	
//	year, month, day := date.Date()
//	
//	newDate := strconv.Itoa(year)+"/"+strconv.Itoa(int(month))+"/"+strconv.Itoa(day)
	
	const layout = "01/02/2006"
	
	newDate := date.Format(layout)
	fmt.Println(newDate)
	return newDate
}
