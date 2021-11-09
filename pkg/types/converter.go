package types

import (
	"jasy/goblog/pkg/logger"
	"strconv"
)

// Int64ToString 将 int64 转换为 string
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

// Uint64ToString 将 uint64 转换为 string
func Uint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}

// StringToInt 将字符串转换为 int
func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		logger.LogError(err)
	}
	return i
}

// StringToUint64 将字符串转换为 uint64
func StringToUint64(str string) uint64 {
	i, err := strconv.Atoi(str)
	if err != nil {
		logger.LogError(err)
	}
	return uint64(i)
}