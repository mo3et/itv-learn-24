package codeq

import "fmt"

// 定义函数 PrintInfo
func PrintInfo(input interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", input, input)
}

func Q2() {
	// 测试 PrintInfo 函数，传入不同类型的参数
	PrintInfo(42)
	PrintInfo("Hello nil interface.")
	PrintInfo(3.14)
	PrintInfo(make([]int, 0, 10))
	PrintInfo(make(map[int]bool, 100))
}
