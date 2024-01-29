package helper

import (
	"crypto/md5"
	"fmt"
	"time"
)

func GetMD5String(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func LoadLocationJakarta() *time.Location {
	jakartaLocation, _ := time.LoadLocation("Asia/Jakarta")
	return jakartaLocation
}

func GetWeekdays(req time.Time) int {
	wd := req.Weekday()
	switch req.Weekday() {
	case 1, 2, 3, 4, 5, 6:
		return int(wd + 1)
	default:
		return 0
	}
}
