package main

import "fmt"

func main() {
	m := make(map[int]string)

	m[0] = "zero"
	m[1] = "one"
	m[2] = "two"

	keys := make([]int, len(m))
	index := 0
	for k := range m {
		keys[index] = k
		index++
		fmt.Println("", k)
	}
}
