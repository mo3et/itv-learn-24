package main

import "fmt"

func cal(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

// defer 在声明时会把函数参数给确定下来，而函数内部的变量会等到调用时获取
func main() {
	a := 1
	b := 2

	defer func() {
		cal("1", a, cal("10", a, b))
	}()
	defer func() {
		cal("2", a, cal("20", a, b))
	}()

	b = 1
}
