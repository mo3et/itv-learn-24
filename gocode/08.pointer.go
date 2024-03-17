package main

import "fmt"

/*
golang 提供了数据结构指针的能力，但是，并不能进行指针运算
*/
func main() {
	i, j := 42, 2701

	p := &i         // 指向i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}
