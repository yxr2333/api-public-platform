package utils

import (
	"crypto/rand"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func GenerateAPIToken(length int) (string, error) {
	token := make([]byte, length)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(token), nil
}

// GeneratePasswordHash 接受一个明文字符串密码，返回加密后的密码和错误（如果存在的话）
func GeneratePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash 接受一个明文密码和一个加密后的密码，返回两者是否匹配的布尔值和错误（如果存在的话）
func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		// 如果有错误，那么密码不匹配
		return false, err
	}

	// 如果没有错误，那么密码匹配
	return true, nil
}
