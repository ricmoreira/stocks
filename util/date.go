package util

import (
	"strings"
	"time"
)

// ParseDate returns a time.Time object for string dates formated as YYYY-MM-DD (eg.: "2018-11-30")
func ParseDate(date string) (time.Time, error) {
	form := "2006-01-02"
	return time.Parse(form, date)
}

// ParseDateTime returns a time.Time object for string dates formated as YYYY-MM-DD HH-MM-SS (eg.: "2018-02-01T10:41:11")
func ParseDateTime(dateTime string) (time.Time, error) {
	t := strings.TrimSpace(dateTime)
	formSeconds := "2006-01-02T15:04:05"
	formMillSec := "2006-01-02T15:04:05.000Z"
	if len(dateTime) == 24 {
		return time.Parse(formMillSec, t)
	}
	return time.Parse(formSeconds, t)
}
