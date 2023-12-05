package helper

import (
	"crypto/md5"
	"fmt"
)

func GetMD5String(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
