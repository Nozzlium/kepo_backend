package helper

import (
	"time"
)

func TimeToString(dateTime time.Time) string {
	current := time.Now()
	if current.Year() == dateTime.Year() {
		if dateTime.Day() == current.Day() &&
			dateTime.Month() == current.Month() {
			return dateTime.Format("15.04")
		}
		return dateTime.Format("02 Jan")
	} else {
		return dateTime.Format("02 Jan 2006")
	}
}
