package utils

import (
	"fmt"
	"testing"
)

func Test_StringToInt(t *testing.T) {
	str := "213"
	num := StringToInt(str)

	fmt.Println("=Utils=Test_StringToInt=1=结果=", str, num)
}
