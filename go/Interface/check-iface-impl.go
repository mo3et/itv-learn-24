package main

import (
	"fmt"
	"reflect"
)

// 断言 Interface implement
// link: https://loesspie.com/2021/11/09/go-assert-type-implement-interface/

type I interface {
	Foo()
}

type T struct{}
type (
	F struct{}
	R string
)

func (t T) Foo() {} // 如缺少，报错 不能将 (*T)(nil) 用作 I 类型，*T 没有实现 I.

// 判断*T是否实现了 I 接口，未实现编译时将报错
// var _ Interface =(*struct)(nil)
var (
	_ I = (*T)(nil) // 将nil转为*T类型后判断是否实现I接口
	_ I = &T{}
	_ I = new(T)
)

// 判断 T 是否实现 I 接口，未实现编译时将报错
var _ I = T{}

func CheckInterfaceImplement() {
	// _ = []int([]int64{}) // cannot convert []int64{} (value of type []int64) to type []int

	// 可被定义为nil的类型，将nil转换为这些类型
	n1 := (*struct{})(nil) // 指针
	fmt.Printf("(*struct{})(nil) nil: %t ,type is: %s\n", n1 == nil, reflect.TypeOf(n1).String())

	n2 := []int(nil) // 切片
	fmt.Printf("[]int(nil) nil: %t ,type is: %s\n", n2 == nil, reflect.TypeOf(n2).String())

	n3 := map[int]bool(nil) // map
	fmt.Printf("map[int]bool(nil) nil: %t ,type is: %s\n", n3 == nil, reflect.TypeOf(n3).String())

	n4 := chan string(nil) // channel
	fmt.Printf("chan string(nil) nil: %t ,type is: %s\n", n4 == nil, reflect.TypeOf(n4).String())

	n5 := (func())(nil) // 函数
	fmt.Printf("(func())(nil) nil: %t ,type is: %s\n", n5 == nil, reflect.TypeOf(n5).String())

	n6 := interface{}(nil) // 接口，可任意赋值
	fmt.Printf("interface{}(nil) nil: %t \n", n6 == nil)

	// 等价于下面
	var _ *struct{} = nil
	var _ []int = nil
	var _ map[int]bool = nil
	var _ chan string = nil
	var _ func() = nil
	var _ interface{} = nil

	// 报错 use of untyped nil
	// var _ = nil
}
