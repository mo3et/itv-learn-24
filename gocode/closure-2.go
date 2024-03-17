package main

import "fmt"

// add 是一个普通函数，接收两个参数并返回和
func add(a, b int) int {
	return a + b
}

// curryAdd 函数是一个闭包，用于实现柯里化
func curryAdd(a int) func(int) int {
	// 在闭包内部定义一个匿名函数，接收一个参数并返回和
	return func(b int) int {
		return add(a, b)
	}
}

func main() {
	//  使用柯里化函数创建新的函数
	add5 := curryAdd(5)

	// 调用新函数，传入参数并获取结果
	res := add5(10)

	fmt.Println(res)
}
