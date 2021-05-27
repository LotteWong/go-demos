package utils

import "crypto/sha1"

func ToSha1(plainText string) string {
	cipherText := string(sha1.New().Sum([]byte(plainText)))
	return cipherText
}
