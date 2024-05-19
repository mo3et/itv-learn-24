package main

import "fmt"

func closu() func() {
	sum := 0
	return func() {
		fmt.Println(sum)
		sum++
	}
}

func main() {
	closurePlus := closu()
	closurePlus()
	closurePlus()
	closurePlus()
	closurePlus()
	closurePlus()
}
