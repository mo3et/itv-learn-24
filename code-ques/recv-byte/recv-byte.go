package main

import "fmt"

func main() {
	str := "reverse string"
	Reverse([]byte(str))
	fmt.Println(str)
}

func Reverse(s []byte) {
	r := len(s) - 1
	l := 0

	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
	fmt.Println(string(s))
}

func Bubble(in []int) []int {
	for i := 0; i < len(in); i++ {
		flag := false
		for j := 0; j < len(in)-i-1; j++ {
			if in[j] > in[j+1] {
				in[j], in[j+1] = in[j+1], in[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return in
}
