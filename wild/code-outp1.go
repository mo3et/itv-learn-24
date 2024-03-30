package main

import (
	"fmt"
	"log"
)

func main() {
	var a int8 = -1
	var b int8 = -128 / a
	fmt.Println(b)

	ma := make([]int, 10000)
	mb := make([]int, 1000)
	if ma[9999] == mb[0] {
		log.Println("nil")
	} else {
		fmt.Println(ma, "\n", mb)
	}
}
