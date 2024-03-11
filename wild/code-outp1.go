package main

import ("fmt")

func main() {
	var a int8 = -1
	var b int8 = -128 / a
	fmt.Println(b)

	ma:=make([]int,10000)
	mb:=make([]int,1000)
}

