package main

import "fmt"

func main() {
	str := "abcdefg1"

	res, _ := reverString(str)
	res2, _ := reverStringez(str)
	fmt.Println(res)
	fmt.Println(res2)
}

func reverString(str string) (string, bool) {
	b := []byte(str)
	if len(b) > 5000 {
		return str, false
	}
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return string(b), true
}

func reverStringez(str string) (string, bool) {
	b := []byte(str)
	if len(b) > 5000 {
		return str, false
	}
	for i, l := 0, len(b)-1; i < l; i, l = i+1, l-1 {
		b[i], b[l] = b[l], b[i]
	}
	return string(b), true
}
