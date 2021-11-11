package password

import (
	"golang.org/x/crypto/bcrypt"
	"jasy/goblog/pkg/logger"
)

func Hash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.LogError(err)
	return string(bytes)
}

func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	logger.LogError(err)
	return err == nil
}

func IsHashed(str string) bool {
	return len(str) == 60
}