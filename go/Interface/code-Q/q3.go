package codeq

import "fmt"

// 定义接口 Shape
type Shape interface {
	Area() float64
}

// 定义结构体 Circle 和 Rectangle，实现接口 Shape
type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (c *Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

func Q3() {
	// 创建包含多个不同形状的切片
	shapes := []Shape{
		&Circle{Radius: 2.5},
		&Rectangle{Width: 3, Height: 4},
		&Circle{Radius: 4},
	}

	// 计算切片中形状的总面积
	totalArea := 0.0
	for _, shape := range shapes {
		totalArea += shape.Area()
		fmt.Println("current Area:", shape.Area())
	}
	fmt.Println("Total Area:", totalArea)
}
