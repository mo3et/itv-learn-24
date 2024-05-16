package main

import (
	"fmt"
	"testing"
)

func Test_ReconvStr(t *testing.T) {
	str := "strajnjkan"

	a, _ := reverString(str)
	fmt.Println(a)
}
