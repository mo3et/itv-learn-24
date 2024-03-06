package simpleimpl

import "fmt"

// 简单用法

// 定义接口 在接口里面定义方法
type Animal interface {
	// 该方法不需要参数，返回值类型为String
	Move() string
}

// 定义一个结构体类型"Dog"
type Dog struct{}

// 实现 Animal 里面的方法 Move，即为实现 `Animal` 接口。
func (d Dog) Move() string {
	return "Dog is Moving."
}

// 创建一个 `Animal`类型的变量，并赋值为一个Dog类型的变量

func AniImplement() {
	// 等同于 aniImpl := Dog{}
	var aniImpl Animal = Dog{} // 使用结构体初始化接口
	fmt.Println(aniImpl.Move())
}
