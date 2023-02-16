package utils

import "strconv"

// StringToInt 字符串-转-int
func StringToInt(str string) (err error, out int) {
	out, err = strconv.Atoi(str)
	return err, out
}

// StringToInt64 字符串-转-int64
func StringToInt64(str string) (err error, out int64) {
	out, err = strconv.ParseInt(str, 10, 64)
	return err, out
}
