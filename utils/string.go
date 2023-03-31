package utils

import "strconv"

// StringToInt 字符串-转-int
func StringToInt(str string) (out int) {
	out, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return out
}

// StringToInt64 字符串-转-int64
func StringToInt64(str string) (out int64) {
	out, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}

	return out
}

// StringToFloat64 字符串-转-float64
func StringToFloat64(str string) (out float64) {
	out, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0.00
	}

	return out
}

