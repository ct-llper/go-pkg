package utils

import "strconv"

func Float32ToString(f float64) (out string) {
	out = strconv.FormatFloat(f, 'E', -1, 32)
	return out
}
