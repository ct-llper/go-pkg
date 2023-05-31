package idcard

import (
	"fmt"
	"testing"
)

var (
	birth1 = "2001-01-01"
	year   = "2001"
	month  = "01"
	day    = "01"
)

func Test_GetAgeReal(t *testing.T) {
	age := GetAgeReal("2020", "06", "01")
	fmt.Println(age)
}
