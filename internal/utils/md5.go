package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(content string) string {
	md5er := md5.New()
	md5er.Write([]byte(content))
	md5Res := md5er.Sum(nil)
	return hex.EncodeToString(md5Res)
}
