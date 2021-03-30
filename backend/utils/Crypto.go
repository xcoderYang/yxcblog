package utils

import "crypto/sha256"

func Sha256(password string)string{
	ans := sha256.Sum256([]byte(password))
	return string(ans[:])
}