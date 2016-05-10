package util

import (
	"time"
)

func GetDate(date time.Time) string {

	const layout = "01/02/2006"

	newDate := date.Format(layout)
	return newDate
}
