package utils

import (
	"os"
	"strings"
)

func DirName(path string) string {
	if len(path) < 1 {
		return path
	}
	pathRune := []rune(path)
	if os.IsPathSeparator(uint8(pathRune[len(pathRune)-1])) {
		pathRune = pathRune[len(pathRune)-1:]
	}
	path = string(pathRune)
	tmp := strings.Split(path, string(os.PathSeparator))
	newPath := strings.Join(tmp[:len(tmp)-1], string(os.PathSeparator))
	return newPath
}
