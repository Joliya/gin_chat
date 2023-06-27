package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// 加盐，提升密码的安全性
func MakePassword(plainpwd, salt string) string {
	return Md5Encode(plainpwd + salt)
}

func ValidPassword(plainpwd string, salt string, password string) bool {
	return Md5Encode(plainpwd+salt) == password
}
