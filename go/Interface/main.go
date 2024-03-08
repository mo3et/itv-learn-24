package main

import (
	"fmt"
	"log"

	codeq "github.com/mo3et/itv-learn-24/go/Interface/code-Q"
)

type (
	Dog    struct{}
	Rabbit struct {
		name   string
		number int
	}
)

func (d *Dog) Moved() string {
	return "Dog Moved."
}

func (d Dog) Speak() {
	fmt.Println("Fuxk you")
}

func (r *Rabbit) Moved() string {
	// name := r.name
	r.name = "no true name"

	return r.name
}

func (r *Rabbit) Speak() {
	fmt.Println("nope")
}

type Animal interface {
	Moved() string
	Speak()
}

func behavior(Ani Animal) {
	Ani.Moved()
	Ani.Speak()
}

func main() {
	// 判断*Rabbit是否实现了Animal接口，未实现 编译时将报错
	var _ Animal = (*Rabbit)(nil)
	var _ Animal = (*Dog)(nil)

	// 实例化
	var aniD Animal = &Dog{}
	var ani Animal = new(Rabbit)

	aniD.Moved()
	ani.Speak()

	// parms := []interface{}{}
	animals := []Animal{&Dog{}, &Rabbit{name: "kitty", number: 6}}
	for _, a := range animals {
		behavior(a)
	}
	fmt.Println()
	for _, a := range animals {
		if _, ok := a.(interface{}); ok {
			log.Print("success.")
		}
	}
	// var d Animal = &Dog{}
	// rimpl := &Rabbit{}
	// var r *Rabbit = &Rabbit{}
	// var a Animal = &Rabbit{}
	// d.Moved()
	// d.Speak()
	// r.Moved()
	// a.Moved()
	// rimpl.Moved()

	codeq.Q1()
	codeq.Q2()
	codeq.Q3()
	codeq.Q4()
	// codeq.Q5()
	// codeq.Q6()
}
