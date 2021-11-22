package test

import (
	"fmt"
	"testing"
)

func TestCompareStr(t *testing.T) {
	a := "abcde"
	b := "abcdf"
	if b > a {
		fmt.Printf("!\n")
	}
}
