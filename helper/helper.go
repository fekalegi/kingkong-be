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
